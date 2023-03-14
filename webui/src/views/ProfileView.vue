<script>
import UploadPhoto from '../components/UploadPhoto.vue';
export default {
    components: {
        UploadPhoto,
    },
    data() {
        return {
            errorMsg: "",
            // Profile data
            userID: parseInt(this.$route.params.userID),
            username: "",
            followersCount: 0,
            followingsCount: 0,
            postsCount: 0,
            bio: "This user have notighing to say",
            proPic64: "",
            isFollowed: false,

            isOwner: false,

            // Edting propic
            showEditPropic: false,
            isEditingPropic: false,

            // Buttons Text
            followTextButton: "Follow",

            // Other Data
            textCounter: 0,
            profilesArray: [],
            textHeader: "",
            typeList: "",

            // Posts data
            posts: [],
            postOffset: 0,
            postLimit: 20,
            showPost: false,
            postViewData: {},

            // Load more data
            busy: false,
            dataAvaible: true,

            // Follower data
            dataGetter: () => { },
            showList: false,

            // Options data
            showOptions: false,
        }
    },
    methods: {
        async getProfile() {
            try {
                let response = await this.$axios.get(`profiles/${this.userID}`, { headers: { 'Authorization': `${localStorage.token}` } })
                this.userID = response.data.user.userID;
                this.username = response.data.user.username;
                this.followersCount = response.data.followersCount;
                this.followingsCount = response.data.followingsCount;
                this.postsCount = response.data.postsCount;
                if (response.data.bio != "")
                    this.bio = response.data.bio;
                this.textCounter = this.bio.length;
                this.proPic64 = response.data.user.userPropic64;
                this.isFollowed = response.data.isFollowed;
                this.followTextButton = this.isFollowed ? "Unfollow" : "Follow";
                this.isOwner = localStorage.userID == this.userID;
            } catch (e) {
                this.errorMsg = e.toString();
            }
        },
        async getPosts() {
            try {
                let response = await this.$axios.get(`/profiles/${this.userID}/posts?limit=${this.postLimit}&offset=${this.postOffset}`, { headers: { 'Authorization': `${localStorage.token}` } });
                if (response.data == null) {
                    this.dataAvaible = false;
                    return;
                }
                this.posts.push(...response.data);
            } catch (e) {
                this.errorMsg = e.toString();
            };
        },
        editingBio() {
            if (this.isOwner) {
                document.getElementsByClassName("top-body-profile-bio-text-counter")[0].style.color = "rgb(0,0,0,0.5)";
                document.getElementsByClassName("top-body-profile-bio-text")[0].style.outline = "auto";
                document.getElementsByClassName("top-body-profile-bio-text")[0].style.outlineColor = "#03C988";

            }
        },
        async saveChangeBio() {
            if (this.isOwner) {
                if (this.bio == "") {
                    this.bio = "This user have notighing to say";
                }
                document.getElementsByClassName("top-body-profile-bio-text")[0].style.outline = "none";
                try {
                    let _ = await this.$axios.put(`/profiles/${this.userID}/bio`, { bio: this.bio }, { headers: { 'Authorization': `${localStorage.token}` } });
                } catch (e) {
                    this.errorMsg = e.toString();
                }
                document.getElementsByClassName("top-body-profile-bio-text-counter")[0].style.color = "#fff";
            }
        },
        editingUsername() {
            if (this.isOwner) {
                document.getElementsByClassName("top-body-profile-username")[0].style.outline = "auto";
                document.getElementsByClassName("top-body-profile-username")[0].style.outlineColor = "#03C988";
            }
        },
        async saveChangeUsername() {
            if (this.isOwner) {
                document.getElementsByClassName("top-body-profile-username")[0].style.outline = "none";
                if (this.username == "" | this.username.length < 3) {
                    this.username = localStorage.username;
                    return
                }
                try {
                    let _ = await this.$axios.put(`/profiles/${this.userID}/username`, { username: this.username }, { headers: { 'Authorization': `${localStorage.token}` } });
                    localStorage.username = this.username;
                } catch (e) {
                    this.errorMsg = e.toString();
                    this.username = localStorage.username;
                }
            }
        },
        getFollowers() {
            this.showList = true;
            this.textHeader = "Followers";
            this.typeList = "simple";
            this.dataGetter = async (profilesArray, limit, offset, dataAvaible) => {
                try {
                    let response = await this.$axios.get(`/profiles/${this.userID}/followers?limit=${limit}&offset=${offset}`, { headers: { 'Authorization': `${localStorage.token}` } });
                    if (response.data == null) {
                        dataAvaible = false;
                        return;
                    }
                    profilesArray.push(...response.data);
                } catch (e) {
                    this.errorMsg = e.toString();
                }
            }
        },
        getFollowings() {
            this.showList = true;
            this.textHeader = "Followings";
            this.typeList = "simple";
            this.dataGetter = async (profilesArray, limit, offset, dataAvaible) => {
                try {
                    let response = await this.$axios.get(`/profiles/${this.userID}/followings?limit=${limit}&offset=${offset}`, { headers: { 'Authorization': `${localStorage.token}` } });
                    if (response.data == null) {
                        dataAvaible = false;
                        return;
                    }
                    profilesArray.push(...response.data);
                } catch (e) {
                    this.errorMsg = e.toString();
                }
            }
        },
        freeLists() {
            this.showList = false;
            this.profilesArray = [];
            this.textHeader = "";
        },
        async follow() {
            if (this.isFollowed) {
                try {
                    let _ = await this.$axios.delete(`profiles/${localStorage.userID}/followings/${this.userID}`, { headers: { 'Authorization': `${localStorage.token}` } });
                    this.isFollowed = false;
                    this.followTextButton = "Follow";
                    this.followersCount--;
                } catch (e) {
                    this.errorMsg = e.toString();
                }
            } else {
                try {
                    let _ = await this.$axios.put(`profiles/${localStorage.userID}/followings/${this.userID}`, {}, { headers: { 'Authorization': `${localStorage.token}` } });
                    this.isFollowed = true;
                    this.followTextButton = "Unfollow";
                    this.followersCount++;
                } catch (e) {
                    this.errorMsg = e.toString();
                }
            }
        },
        openPost(post) {
            this.showPost = true;
            this.postViewData = post;
        },
        exitPost() {
            this.showPost = false;
            this.postViewData = {};
        },
        updateLike(data) {
            this.posts.forEach(post => {
                if (post.postID == data.postID) {
                    post.liked = data.liked;
                    post.likesCount = data.liked ? post.likesCount + 1 : post.likesCount - 1;
                }
            });
        },
        loadMoreContents() {
            console.log("ciaos");
            if (this.busy || !this.dataAvaible) return;
            this.busy = true;
            this.postOffset += this.postLimit;
            this.getPosts();
            this.busy = false;
        },
        doLogout() {
            localStorage.clear();
            this.$router.push('/login');
        },
        getBans() {
            this.showList = true;
            this.textHeader = "Bans";
            this.typeList = "ban";
            this.dataGetter = async (profilesArray, limit, offset, dataAvaible) => {
                try {
                    let response = await this.$axios.get(`/profiles/${this.userID}/bans?limit=${limit}&offset=${offset}`, { headers: { 'Authorization': `${localStorage.token}` } });
                    if (response.data == null) {
                        dataAvaible = false;
                        return;
                    }
                    profilesArray.push(...response.data);
                } catch (e) {
                    this.errorMsg = e.toString();
                }
            }
        },
        closeOptions() {
            setTimeout(() => {
                this.showOptions = false;
            }, 500);
        },
        async banUser() {
            try {
                let _ = await this.$axios.put(`/profiles/${localStorage.userID}/bans/${this.userID}`, {}, { headers: { 'Authorization': `${localStorage.token}` } });
                this.$router.push(`/profiles/${localStorage.userID}`);
            } catch (e) {
                this.errorMsg = e.toString();
            }
            this.showOptions = false;
        },
        async deletePost(postID) {
            try {
                let _ = await this.$axios.delete(`profiles/${localStorage.userID}/posts/${postID}`, { headers: { 'Authorization': `${localStorage.token}` } });
                this.posts = this.posts.filter(post => post.postID != postID);
                this.exitPost();
            } catch (e) {
                this.errorMsg = e.toString();
            }
        },
        updateProfile() {
            this.getProfile();
            localStorage.propic64 = this.proPic64;
        },
        async deleteProfile() {
            try {
                let _ = await this.$axios.delete(`profiles/${localStorage.userID}`, { headers: { 'Authorization': `${localStorage.token}` } });
                this.doLogout();
            } catch (e) {
                this.errorMsg = e.toString();
            }
        },
        async resetProfilePic() {
            try {
                let response = await this.$axios.put(`profiles/${localStorage.userID}/reset-profile-picture`, {}, { headers: { 'Authorization': `${localStorage.token}` } });
                this.proPic64 = response.data.userPropic64;
                this.updateProfile();
            } catch (e) {
                this.errorMsg = e.toString();
            }
        }
    },
    beforeMount() {
        if (!localStorage.token) {
            this.$router.push('/login');
        }
        if (localStorage.userID === this.$route.params.userID) {
            this.isOwner = true;
        }
    },

    mounted() {
        this.getProfile();
        this.getPosts();

        if (this.isOwner) {
            document.getElementsByClassName("top-body-profile-bio-text")[0].style.cursor = "text";
            document.getElementsByClassName("top-body-profile-username")[0].style.cursor = "text";

            document.getElementsByClassName("top-profile-picture")[0].style.cursor = "pointer";
        }

        document.addEventListener('scroll', e => {
            if (document.documentElement.scrollTop + window.innerHeight >= document.documentElement.scrollHeight*0.8) {
                this.loadMoreContents();
            }
        });
    },

    beforeRouteUpdate(to, from) {
        this.posts = [];
        this.postOffset = 0;
        this.dataAvaible = true;

        this.userID = parseInt(to.params.userID);
        this.getProfile();
        this.getPosts();
    },
}

</script>


<template>
    <ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"></ErrorMsg>
    <UploadPhoto v-if="isEditingPropic" :photoType="'proPic'" @exit-upload-form="isEditingPropic = false"
        @refresh-data="updateProfile" @error-occured="errorMsg = value"> </UploadPhoto>
    <div class="top-profile-container">
        <div class="top-profile-picture" @mouseover="showEditPropic = isOwner" @mouseleave="showEditPropic = false">
            <div class="edit-propic" v-if="showEditPropic">
                <button class="reset-propic-button">
                    <font-awesome-icon icon="fa-solid fa-xmark" size="lg" color="#fff" @click="resetProfilePic" />
                </button>
                <button class="edit-propic-button" @click="isEditingPropic = true">
                    <font-awesome-icon icon="fa-regular fa-pen-to-square" size="lg" color="#fff" />
                </button>
            </div>
            <img :src="`data:image/jpg;base64,${proPic64}`">
        </div>
        <div class="top-body-profile-container">
            <div class="profile-options-button-container">
                <button class="profile-options-button" @click="showOptions = true" @focusout="closeOptions">
                    <font-awesome-icon icon="fa-solid fa-ellipsis" />
                </button>
                <div v-if="showOptions && isOwner" class="profile-options-menu">
                    <div class="options-menu">
                        <div class="options-menu-item" @click="getBans">
                            <span>Bans list</span>
                        </div>
                        <div class="options-menu-item" @click="deleteProfile">
                            <span>Delete profile</span>
                        </div>
                        <div class="options-menu-item" @click="doLogout">
                            <span>Logout</span>
                        </div>
                    </div>
                </div>
                <div v-else-if="showOptions" class="profile-options-menu">
                    <div class="options-menu-item" @click="banUser">
                        <span>Ban this user</span>
                    </div>
                </div>
            </div>
            <input :readonly="!isOwner" v-model="username" class="top-body-profile-username" @focusin="editingUsername"
                @focusout="saveChangeUsername" @input="checkUsername" maxlength="13" spellcheck="false">
            <div class="top-body-profile-bio-container">
                <span class="top-body-profile-bio-text-counter">{{ textCounter }}/100</span>
                <span class="top-body-profile-bio-label">Bio</span>
                <textarea :readonly="!isOwner" @focusin="editingBio" @focusout="saveChangeBio" @input="countChar"
                    v-model="bio" class="top-body-profile-bio-text" spellcheck="false" maxlength="100" rows="2"></textarea>
            </div>
            <div class="top-body-profile-stats-container">
                <div class="profile-stats" @click="goToPost">
                    <span class="profile-stats-text">posts</span>
                    <span class="profile-stats-number">{{ postsCount }}</span>
                </div>
                <div class="profile-stats" @click="getFollowers">
                    <span class="profile-stats-text">followers</span>
                    <span class="profile-stats-number">{{ followersCount }}</span>
                </div>
                <div class="profile-stats" @click="getFollowings">
                    <span class="profile-stats-text">followings</span>
                    <span class="profile-stats-number">{{ followingsCount }}</span>
                </div>
            </div>
            <div class="top-body-profile-actions" v-if="!isOwner">
                <button class="profile-actions-button follow-button" @click="follow()"> {{ followTextButton }} </button>
            </div>
        </div>
    </div>

    <ProfilesList v-if="showList" :dataGetter="dataGetter" :textHeader="textHeader" :typeList="typeList"
        @exit-list="freeLists" />
        <div v-if="showPost" class="post-view" @click.self="exitPost">
            <Post :postData="postViewData" @delete-post="deletePost" />
        </div>

    <div class="posts-container">
        <span v-if="(posts.length === 0)" class="posts-grid-nopost-text"> There are no posts yet </span>
        <div class="posts-grid-container" v-if="posts.length > 0">
            <div v-for="post in posts" :key="post.postID" class="posts-grid-post" @click="openPost(post)">
                <img :src="`data:image/jpg;base64,${post.image}`" loading="lazy" class="posts-grid-post-image"
                    :id="post.postID">
            </div>
        </div>

    </div>
</template>

