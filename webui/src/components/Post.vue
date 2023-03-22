<script>

import ProfilesList from '@/components/ProfilesList.vue';

export default {
    emits: ['update-like', 'delete-post', 'error-occurred'],
    components: {
        ProfilesList,
    },
    props: {
        postData: { type: Object, required: true },
    },
    data() {
        return {

            // Data User
            ownerID: this.postData.user['userID'],
            username: this.postData.user['username'],
            proPic64: this.postData.user['userPropic64'],

            // Data Post
            postID: this.postData.postID,
            isLiked: this.postData.liked,
            image64: this.postData.image,
            likesCount: this.postData.likesCount,
            commentsCount: this.postData.commentsCount,
            timestamp: this.postData.timestamp,
            captionPost: this.postData.caption != '' ? this.postData.caption : "There is no caption for this post",

            isOwner: false,
            isHoverComment: false,

            profilesList: [],
            textCounter: 0,
            showList: false,
            textHeader: "",
            typeList: "",
            dataGetter: () => { },
            dataUpdater: () => { },
            additionalData: {},

            errorMsg: "",
        };
    },
    methods: {
        since(timestamp) {
            return this.$utils.since(timestamp);
        },
        async like() {
            try {
                let _ = await this.$axios.put(`profiles/${this.ownerID}/posts/${this.postID}/likes/${localStorage.userID}`, {}, { headers: { "Authorization": `${localStorage.token}` } });
                this.isLiked = true;
                this.$emit('update-like', { postID: this.postID, liked: true });
                this.likesCount++;
            } catch (e) {
                this.errorMsg = this.$utils.errorToString(e);
                this.$emit('error-occurred', this.errorMsg);
            }
        },
        async unlike() {
            try {
                let _ = await this.$axios.delete(`profiles/${this.ownerID}/posts/${this.postID}/likes/${localStorage.userID}`, { headers: { "Authorization": `${localStorage.token}` } });
                this.isLiked = false
                this.$emit('update-like', { postID: this.postID, liked: false });
                this.likesCount--;
            } catch (e) {
                this.errorMsg = this.$utils.errorToString(e);
                this.$emit('error-occurred', this.errorMsg);
            }
        },
        toggleComment() {
            this.isHoverComment = !this.isHoverComment;
        },
        getComments() {
            this.showList = true;
            this.textHeader = "Comments";
            this.typeList = "comment";
            this.additionalData = { 'commentsCount': this.commentsCount, 'postID': this.postID, 'ownerID': this.ownerID };
            this.dataGetter = async (profilesArray, limit, offset, dataAvaible) => {
                try {
                    let response = await this.$axios.get(`/profiles/${this.ownerID}/posts/${this.postID}/comments?limit=${limit}&offset=${offset}`, { headers: { 'Authorization': `${localStorage.token}` } });
                    if (response.data == null) {
                        dataAvaible = false;
                        return;
                    }
                    profilesArray.push(...response.data);
                } catch (e) {
                    this.errorMsg = this.$utils.errorToString(e);
                }
            };
            this.dataUpdater = (entries, values) => {
                if (values['opType'] == 'insert') {
                    entries.push(values['value']);
                    this.commentsCount++;
                } else if (values['opType'] == 'delete') {
                    const index = entries.findIndex((entry) => entry.commentID == values['value']);
                    entries.splice(index, 1);
                    this.commentsCount--;
                }
            };
        },
        getLikes() {
            this.showList = true;
            this.textHeader = "Likes";
            this.typeList = "simple";

            this.additionalData = { 'likesCount': this.likesCount };
            this.dataGetter = async (profilesArray, limit, offset, dataAvaible) => {
                try {
                    let response = await this.$axios.get(`/profiles/${this.ownerID}/posts/${this.postID}/likes?limit=${limit}&offset=${offset}`, { headers: { 'Authorization': `${localStorage.token}` } });
                    if (response.data == null) {
                        dataAvaible = false;
                        return;
                    }
                    profilesArray.push(...response.data);
                } catch (e) {
                    this.errorMsg = this.$utils.errorToString(e);
                }
            }
        },
        editingCaption() {
            if (this.isOwner) {
                document.querySelectorAll(".post-tail-caption-text-counter")[0].style.color = "rgb(0,0,0,0.5)";
                document.querySelectorAll(".post-tail-caption")[0].style.outline = "auto";
                document.querySelectorAll(".post-tail-caption")[0].style.outlineColor = "#03C988";
            }
        },
        async saveChangeCaption() {
            if (this.isOwner) {
                if (this.captionPost == "") {
                    this.captionPost = "This user have notighing to say";
                }
                document.querySelectorAll(".post-tail-caption")[0].style.outline = "none";
                try {
                    await this.$axios.put(`/profiles/${this.ownerID}/posts/${this.postID}/caption`, { caption: this.captionPost }, { headers: { "Authorization": `${localStorage.token}` } });
                }
                catch (e) {
                    this.errorMsg = this.$utils.errorToString(e);
                    this.$emit('error-occurred', this.errorMsg);
                }
                document.querySelectorAll(".post-tail-caption-text-counter")[0].style.color = "#fff";
            }
        },
        goToProfile() {
            this.$router.push(`/profiles/${this.ownerID}`);
        },
        deletePost() {
            this.$emit('delete-post', this.postID);
        },
    },
    beforeMount() {
        if (this.ownerID == localStorage.userID) {
            this.isOwner = true;
        }
    },
    mounted() {
    },
    afterMount() {
        if (this.isOwner) {
            document.querySelectorAll(".post-tail-caption-text")[0].style.cursor = "text";
        }
    },

    watch: {
        captionPost() {
            this.textCounter = this.captionPost.length;
            if (this.captionPost.includes("\n")) {
                this.captionPost = this.captionPost.replace("\n", "");
                this.captionPost = this.captionPost.replace("/", "");
                this.captionPost = this.captionPost.replace("\\", "");
            }
        },
    },
}
</script>

<template>
    <div class="post-containter">
        <div class="post-header">
            <img class="post-header-propic-img" :src="`data:image/jpg;base64,${proPic64}`" alt="" loading="lazy">
            <span @click="goToProfile" class="post-header-username">{{ username }}</span>
            <span class="post-header-timestamp">{{ since(timestamp) }}</span>
        </div>
        <img class="post-img" :src="`data:image/jpg;base64,${image64}`" @dblclick="() => isLiked ? unlike() : like()"
            loading="lazy">
        <div class="post-tail">
            <div class="post-tail-like-comment-options">
                <button class="post-tail-button fa-layers fa-fw" name="like" @click.self="isLiked ? unlike() : like()">
                    <font-awesome-icon v-if="!isLiked" icon="fa-regular fa-heart" @click="like" />
                    <font-awesome-icon v-else icon="fa-solid fa-heart" style="color:red" @click="unlike" />
                    <span class="fa-layers-counter bg-secondary h2">{{ likesCount }}</span>
                    <span class="title-button" @click.self="getLikes">
                        Likes
                    </span>
                </button>
                <button class="post-tail-button fa-layers fa-fw" name="comment" @click="getComments">
                    <font-awesome-icon v-if="!isHoverComment" icon="fa-regular fa-comment" @mouseover="toggleComment" />
                    <font-awesome-icon v-else icon="fa-solid fa-comment" @mouseout="toggleComment"
                        style="opacity:0.9" />
                    <span class="fa-layers-counter bg-secondary h2">{{ commentsCount }}</span>
                    <span class="title-button">
                        Comments
                    </span>
                </button>
                <button v-if="isOwner" class="post-tail-delete-button" name="delete" @click="deletePost">
                    <font-awesome-icon icon="fa-regular fa-trash-can" />
                </button>
            </div>
            <textarea :readonly="!isOwner" @focusin="editingCaption" @focusout="saveChangeCaption" v-model="captionPost"
                spellcheck="false" maxlength="64" rows="2" class="post-tail-caption"></textarea>
            <span class="post-tail-caption-text-counter">{{ textCounter }}/64 </span>
        </div>
    </div>
    <ProfilesList @exit-list="showList = false" v-if="showList" :textHeader="textHeader" :dataGetter="dataGetter"
        :dataUpdater="dataUpdater" :typeList="typeList" :argv="additionalData" />
</template>


<style>
.stats-lists-component {
    z-index: 3;
}

.post-containter {
    width: 30em;
    height: auto;

    background-color: rgb(0, 0, 0, 0.1);
    box-shadow: 0 4px 5px -2px rgba(0, 0, 0, 0.4);

    border-radius: 0.5em;

    margin-bottom: 4em;

    display: flex;
    flex-direction: column;
    justify-content: space-between;

    z-index: 2;
}

.post-header {
    width: 100%;
    height: 4em;
    background-color: #fff;
    border-radius: 0.5em 0.5em 0 0;
}

.post-header-username {
    font-size: 1.3em;
    font-weight: 600;


    margin-left: 0.5em;

    position: relative;
    top: 0.4em;

    cursor: pointer;
}

.post-header-propic-img {
    width: 2.5em;
    height: 2.5em;
    border-radius: 10em;

    margin: 0.8em 0em 0em 1em;

    object-fit: cover;
}

.post-header-timestamp {
    font-size: 1em;
    font-weight: 500;

    color: rgb(0, 0, 0, 0.5);

    position: relative;
    top: 0.6em;
    left: 1em;
}

.post-img {
    width: 30em;
    height: 30em;
}

.post-tail {
    width: 100%;
    height: 8em;
    background-color: #fff;
    border-radius: 0 0 0.5em 0.5em;

    padding: 0.2em 0.5em 0 0.5em;
}


.post-tail-like-comment-options {
    width: 100%;
    height: 2em;

    font-size: 1.5em;
    border-bottom: 1px solid rgb(0, 0, 0, 0.2);

    padding: 0 0 1em 0;

    display: flex;
    flex-direction: row;
    align-items: center;
}

.post-tail-button {
    border: none;
    background-color: transparent;
    outline: none;

    margin-right: 1em;

    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;

}

.title-button {
    font-size: 0.5em;
    font-weight: 500;
    color: rgb(0, 0, 0, 0.5);
    cursor: pointer;

    position: relative;
    top: 2em;

}

.post-tail-delete-button {
    margin-left: auto;

    border: none;
    background-color: transparent;
    outline: none;

    font-size: 0.8em;
    margin-top: 0.3em;
}

.post-tail-caption {
    margin: 0.5em 0 0 0;

    font-size: 1em;
    font-weight: 400;

    width: 100%;

    border: none;
    resize: none;

    cursor: default;

    position: relative;
    outline: none;
}

.post-tail-caption-text-counter {
    position: relative;
    color: #fff;
    float: right;
    font-size: 0.7em;
    line-height: 0;
}
</style>