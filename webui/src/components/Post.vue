<script>
import ProfilesList from './ProfilesList.vue';

export default {
    props: {
        userID: { type: Number, required: true },
        postID: { type: String, required: true },
        username: { type: String, required: true },
        caption: { type: String, required: true },
        image: { type: String, required: true },
        timestamp: { type: String, required: true },
        liked: { type: Boolean, required: true },
    },
    data() {
        return {
            isOwner: false,
            isHoverComment: false,
            isLiked: this.liked,
            likesList: [],
            captionPost: "There is no caption for this post",
            textCounter: 0,
        };
    },
    methods: {
        like() {
            this.isLiked = !this.isLiked;
        },
        toggleComment() {
            this.isHoverComment = !this.isHoverComment;
        },
        async getLikes() {
            try {
                let response = await this.$axios.get(`/profiles/${this.userID}/posts/${this.postID}/likes`, { headers: { "Authorization": `${localStorage.token}` } });
                this.likesList = response.data;
            }
            catch (e) {
                this.errormsg = e.data();
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
                    let _ = await this.$axios.put(`/profiles/${this.userID}/posts/${this.postID}/caption`, { caption: this.caption }, { headers: { "Authorization": `${localStorage.token}` } });
                }
                catch (e) {
                    this.errorMsg = e.toString();
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
            this.$router.push(`/profiles/${this.userID}`);
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
        if (this.userID == localStorage.userID) {
            this.isOwner = true;
        }
        if (this.caption != "") {
            this.captionPost = this.caption;
        }
    },
    mounted() {
        if (this.isOwner) {
            document.getElementsByClassName("post-tail-caption-text")[0].style.cursor = "text";
        }
    }
}
</script>

<template>
    <div class="post-containter">
        <div class="post-header">
            <img class="post-header-propic-img" src="https://i.imgur.com/2oUpQtj.jpg" alt="" loading="lazy">
            <span @click="goToProfile" class="post-header-username">{{ username }}</span>
            <span class="post-header-timestamp">{{ calcTimestamp(new Date(timestamp)) }}</span>
        </div>
        <img class="post-img" :src="image" @dblclick="like" loading="lazy">
        <div class="post-tail">
            <div class="post-tail-like-comment-options">
                <button class="post-tail-like-button" name="like">
                    <font-awesome-icon v-if=!isLiked icon="fa-regular fa-heart" @click="like" />
                    <font-awesome-icon v-else icon="fa-solid fa-heart" @click="like" @dblclick="getLikes"
                        style="color:red" />
                </button>
                <button class="post-tail-comment-button" name="comment">
                    <font-awesome-icon v-if="!isHoverComment" icon="fa-regular fa-comment" v-on:mouseover="toggleComment" />
                    <font-awesome-icon v-else icon="fa-solid fa-comment" v-on:mouseout="toggleComment"
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
    <ProfilesList v-if="(likesList.length != 0)"> </ProfilesList>
</template>

<style>
.post-containter {
    width: 30em;
    height: 40em;

    background-color: rgb(0, 0, 0, 0.1);
    box-shadow: 0 4px 5px -2px rgba(0, 0, 0, 0.4);

    border-radius: 0.5em;

    margin-bottom: 6em;

    position: relative;
    left: 0;


    display: flex;
    flex-direction: column;
    justify-content: space-between;

    z-index: 3;
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
}

.post-header-timestamp{
    font-size: 1em;
    font-weight: 500;

    color: rgb(0, 0, 0, 0.5);

    position: relative;
    top: 0.6em;
    left: 1em;
}

.post-img {
    width: 100%;
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
