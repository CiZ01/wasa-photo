<script>

import ProfilesList from './ProfilesList.vue';
import LikeHeader from '../components/LikeHeader.vue';
import CommentHeader from '../components/CommentHeader.vue';
import CommentFooter from '../components/CommentFooter.vue';
import utils from '../services/utils.js';

export default {
    emits: ['update-like', 'delete-post'],
    components: {
        ProfilesList,
        LikeHeader,
        CommentHeader,
        CommentFooter,
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
            captionPost: this.postData.caption ? this.postData.caption : "There is no caption for this post",

            isOwner: false,
            isHoverComment: false,

            profilesList: [],
            textCounter: 0,
            showList: false,
            textHeader: "",
            componentEntries: "",
            customHeader: null,
            customFooter: null,
            typeEntry: '',
            dataGetter: () => { },
            dataUpdater: () => { },
            additionalData: {},

            errorMsg: "",
        };
    },
    methods: {
        since(timestamp) {
            return utils.since(timestamp);
        },
        async like() {
            try {
                let _ = await this.$axios.put(`profiles/${this.ownerID}/posts/${this.postID}/likes/${localStorage.userID}`, {}, { headers: { "Authorization": `${localStorage.token}` } });
                this.isLiked = true;
                this.$emit('update-like', { postID: this.postID, liked: true });

            } catch (e) {
                this.errorMsg = e.toString();
            }
        },
        async unlike() {
            try {
                let _ = await this.$axios.delete(`profiles/${this.ownerID}/posts/${this.postID}/likes/${localStorage.userID}`, { headers: { "Authorization": `${localStorage.token}` } });
                this.isLiked = false
                this.$emit('update-like', { postID: this.postID, liked: false });
            } catch (e) {
                this.errorMsg = e.toString();
            }
        },
        toggleComment() {
            this.isHoverComment = !this.isHoverComment;
        },
        async getComments() {
            this.showList = true;
            this.textHeader = "";
            this.componentEntries = "CommentEntry";
            this.customHeader = CommentHeader;
            this.customFooter = CommentFooter;
            this.additionalData = { 'commentsCount': this.commentsCount, 'postID': this.postID, 'ownerID': this.ownerID};
            this.dataGetter = async (profilesArray, limit, offset, dataAvaible) => {
                try {
                    let response = await this.$axios.get(`/profiles/${this.ownerID}/posts/${this.postID}/comments?limit=${limit}&offset=${offset}`, { headers: { 'Authorization': `${localStorage.token}` } });
                    if (response.data == null) {
                        dataAvaible = false;
                        return;
                    }
                    profilesArray.push(...response.data);
                } catch (e) {
                    this.errorMsg = e.toString();
                }
            };
            this.dataUpdater = (entries, values) => {
                if (values.opType == 'insert'){
                    const value = values.value;
                    entries.push(value);
                }else if (values.opType == 'delete'){
                    const value = values.value;
                    entries = entries.filter(entry => entry.commentID != value);
                }
            };
        },
        async getLikes() {
            this.showList = true;
            this.textHeader = "";
            this.componentEntries = "SimpleProfileEntry";

            this.customFooter = null;
            this.customHeader = LikeHeader;
            this.additionalData = { 'likesCount': this.likesCount };
            this.dataGetter = async (profilesArray, limit, offset, dataAvaible) => {
                try {
                    let response = await this.$axios.get(`/profiles/${this.ownerID}/posts/${this.postID}/likes?limit=${limit}&offset=${offset}`, { headers: { 'Authorization': `${localStorage.token}` } });
                    if (response.data == null) {
                        dataAvaible = false;
                        return;
                    }
                    profilesArray.push(...response.data);
                    console.log(profilesArray);
                } catch (e) {
                    this.errorMsg = e.toString();
                }
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
}
</script>

<template>
    <ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''" />
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
                <button class="post-tail-button" name="like" @click="isLiked ? unlike() : like()">
                    <font-awesome-icon v-if=(!isLiked) icon="fa-regular fa-heart" />
                    <font-awesome-icon v-else icon="fa-solid fa-heart" style="color:red" />
                    <span class="title-button" @click.self="getLikes">
                        Likes
                    </span>
                </button>
                <button class="post-tail-button" name="comment" @click="getComments">
                    <font-awesome-icon v-if="!isHoverComment" icon="fa-regular fa-comment" v-on:mouseover="toggleComment" />
                    <font-awesome-icon v-else icon="fa-solid fa-comment" v-on:mouseout="toggleComment"
                        style="opacity:0.9" />
                    <span class="title-button">
                        Comments
                    </span>
                </button>
                <button v-if="isOwner" class="post-tail-delete-button" name="delete" @click="deletePost">
                    <font-awesome-icon icon="fa-regular fa-trash-can" />
                </button>
            </div>
            <textarea :readonly="!isOwner" @focusin="editingCaption" @focusout="saveChangeCaption" @input="countChar"
                v-model="captionPost" spellcheck="false" maxlength="64" rows="2" class="post-tail-caption"></textarea>
            <span class="post-tail-caption-text-counter">{{ textCounter }}/64 </span>
        </div>
    </div>
    <ProfilesList @exit-list="showList = false" class="stats-lists-component" v-if="showList"
        :componentHeader="customHeader" :textHeader="textHeader" :dataGetter="dataGetter" :dataUpdater="dataUpdater"
        :componentEntries="componentEntries" :componentFooter="customFooter" :argv="additionalData" />
</template>


<style>
.stats-lists-component {
    z-index: 3;
}

.post-tail-button {
    border: none;
    background-color: transparent;
    outline: none;
    margin-right: 0.5em;

    display: flex;
    flex-direction: column;
    align-items: center;

}

.post-containter {
    width: 30em;
    height: auto;

    background-color: rgb(0, 0, 0, 0.1);
    box-shadow: 0 4px 5px -2px rgba(0, 0, 0, 0.4);

    border-radius: 0.5em;


    position: relative;
    left: 0;

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

    display: flex;
    flex-direction: row;
    align-items: center;
}

.title-button {
    font-size: 0.5em;
    font-weight: 500;
    color: rgb(0, 0, 0, 0.5);
    cursor: pointer;

    position: relative;

}

.post-tail-delete-button {
    float: right;

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