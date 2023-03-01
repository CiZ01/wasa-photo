<script>
import ProfilesList from './ProfilesList.vue';
import LikeHeader from '../components/LikeHeader.vue';
import CommentHeader from '../components/CommentHeader.vue';

export default {
    emits: ['update-like', 'toogle-navbar'],
    components: {
        ProfilesList,
        LikeHeader,
        CommentHeader,
    },
    props: {
        owner: { type: Object, required: true },
        postID: { type: String, required: true },
        caption: { type: String, required: true },
        image: { type: String, required: true },
        timestamp: { type: String, required: true },
        liked: { type: Boolean, required: true },
    },
    data() {
        return {

            // Data User
            ownerID: this.$props.owner['userID'],
            username: this.$props.owner['username'],
            proPic64: this.$props.owner['userPropic64'],


            isOwner: false,
            isHoverComment: false,
            isLiked: this.liked,
            profilesList: [],
            captionPost: "There is no caption for this post",
            textCounter: 0,
            errorMsg: "",
            showList: false,
            customHeader: "",
            titleHeaderList: "",
        };
    },
    methods: {
        async like() {
            try {
                let _ = await this.$axios.put(`/profiles/${this.ownerID}/posts/${this.postID}/likes/${localStorage.userID}`, {}, { headers: { "Authorization": `${localStorage.token}` } });
                this.isLiked = true;
                this.$emit('update-like', { postID: this.postID, liked: true });

            } catch (e) {
                localStorage.errorMessage = e.response.data;
            }
        },
        async unlike() {
            try {
                let _ = await this.$axios.delete(`/profiles/${this.ownerID}/posts/${this.postID}/like/${localStorage.userID}`, { headers: { "Authorization": `${localStorage.token}` } });
                this.isLiked = false
                this.$emit('update-like', { postID: this.postID, liked: false });
            } catch (e) {
                localStorage.errorMessage = e.response.data;
            }
        },
        toggleComment() {
            this.isHoverComment = !this.isHoverComment;
        },
        async getComments() {
            try {
                let response = await this.$axios.get(`/profiles/${this.ownerID}/posts/${this.postID}/comments`, { headers: { "Authorization": `${localStorage.token}` } });
                this.profilesList = response.data;
                this.customHeader = CommentHeader;
                this.titleHeaderList = "Comments";
                this.showList = true;
                console.log(this.profilesList);

            } catch (e) {
                localStorage.errorMessage = e.response.data;
            }
        },
        async getLikes() {
            try {
                let response = await this.$axios.get(`/profiles/${this.ownerID}/posts/${this.postID}/likes`, { headers: { "Authorization": `${localStorage.token}` } });
                this.profilesList = response.data;
                this.customHeader = LikeHeader;
                this.titleHeaderList = "Likes";
                this.showList = true;
            }
            catch (e) {
                localStorage.errorMessage = e.response.data;
            }
        },
        editingCaption() {
            if (this.isOwner) {
                document.getElementsByClassName("post-tail-caption-text-counter")[0].style.color = "rgb(0,0,0,0.5)";
                document.getElementsByClassName("post-tail-caption")[0].style.outline = "auto";
            }
        },
        async saveChangeCaption() {
            if (this.isOwner) {
                if (this.bio == "") {
                    this.bio = "This user have notighing to say";
                }
                document.getElementsByClassName("post-tail-caption")[0].style.outline = "none";
                try {
                    let _ = await this.$axios.put(`/profiles/${this.ownerID}/posts/${this.postID}/caption`, { caption: this.caption }, { headers: { "Authorization": `${localStorage.token}` } });
                }
                catch (e) {
                    localStorage.errorMessage = e.response.data;
                }
                document.getElementsByClassName("post-tail-caption-text-counter")[0].style.color = "#fff";
            }
        },
        countChar() {
            this.textCounter = this.caption.length;
            if (this.caption.includes("\n")) {
                this.caption = this.caption.replace("\n", "");
            }
        },
        goToProfile() {
            this.$router.push(`/profiles/${this.ownerID}`);
        },
        calcTimestamp(timestamp) {
            const ONE_MINUTE = 60 * 1000;
            const ONE_HOUR = 60 * ONE_MINUTE;
            const ONE_DAY = 24 * ONE_HOUR;
            const ONE_WEEK = 7 * ONE_DAY;

            const diff = Math.abs(Date.now() - timestamp);
            const elapsedMinutes = Math.floor(diff / ONE_MINUTE);
            const elapsedHours = Math.floor(diff / ONE_HOUR);
            const elapsedDays = Math.floor(diff / ONE_DAY);
            const elapsedWeeks = Math.floor(diff / ONE_WEEK);

            if (elapsedWeeks > 0) {
                return `${elapsedWeeks}w`;
            } else if (elapsedDays > 0) {
                return `${elapsedDays}d`;
            } else if (elapsedHours > 0) {
                return `${elapsedHours}h`;
            } else {
                return `${elapsedMinutes}m`;
            }
        },
    },
    beforeMount() {
        if (this.ownerID == localStorage.userID) {
            this.isOwner = true;
        }
        if (this.caption != "") {
            this.captionPost = this.caption;
        }
    },
    mounted() {

    },
    afterMount() {
        if (this.isOwner) {
            document.getElementsByClassName("post-tail-caption-text")[0].style.cursor = "text";
        }
    },
    watch: {
        errorMsg: function () {
            this.$emit("errorOccurred", this.errorMsg);
            console.log(this.errorMsg);
        },
    },
}
</script>

<template>
    <div class="post-containter">
        <div class="post-header">
            <img class="post-header-propic-img" :src="`data:image/jpg;base64,${proPic64}`" alt="" loading="lazy">
            <span @click="goToProfile" class="post-header-username">{{ username }}</span>
            <span class="post-header-timestamp">{{ calcTimestamp(new Date(timestamp)) }}</span>
        </div>
        <img class="post-img" :src="`data:image/jpg;base64,${image}`" @dblclick="() => isLiked ? unlike() : like()"
            loading="lazy">
        <div class="post-tail">
            <div class="post-tail-like-comment-options">
                <button class="post-tail-like-button" name="like">
                    <font-awesome-icon v-if=(!isLiked) icon="fa-regular fa-heart" @click="like" />
                    <font-awesome-icon v-else icon="fa-solid fa-heart" @click="unlike" style="color:red" />
                </button>
                <button class="post-tail-comment-button" name="comment">
                    <font-awesome-icon v-if="!isHoverComment" icon="fa-regular fa-comment" v-on:mouseover="toggleComment" />
                    <font-awesome-icon v-else icon="fa-solid fa-comment" @click="getComments" v-on:mouseout="toggleComment"
                        style="opacity:0.9" />
                </button>
                <button v-if="isOwner" class="post-tail-options-button" name="options">
                    <font-awesome-icon icon="fa-solid fa-ellipsis" />
                </button>
            </div>
            <textarea :readonly="!isOwner" @focusin="editingCaption" @focusout="saveChangeCaption" @input="countChar"
                v-model="captionPost" spellcheck="false" maxlength="64" rows="2" class="post-tail-caption"></textarea>
            <span class="post-tail-caption-text-counter">{{ textCounter }}/64 </span>
        </div>
    </div>
    <ProfilesList @exitList="() => !showList" class="stats-lists-component" v-if="showList" :profiles="profilesList"
        :componentHeader="customHeader" :textHeader="titleHeaderList"> </ProfilesList>
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


    position: relative;
    left: 0;

    margin-bottom: 6em;

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
}

.post-tail-like-comment-options>button {
    border: none;
    background-color: transparent;
    outline: none;
    margin-right: 0.5em;
}

.post-tail-options-button {
    float: right;
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