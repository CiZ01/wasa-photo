<script>
import utils from "../services/utils.js";

export default {
    emits: ["exit-list-from-entry", 'error-occurred', 'data-update'],
    props: {
        /*
        {
            commentID,
            postID,
            ownerID,
            user: {
                userID,
                username,
                userPropic64,
            },
            text,
            timestamp
        }
        */
        data: { type: Object, required: true },
    },
    data() {
        return {
            propic64: this.$props.data.user["userPropic64"],
            username: this.$props.data.user["username"],
            userID: this.$props.data.user["userID"],
            
            ownerID: this.$props.data.ownerID,
            postID: this.$props.data.postID,
            text: this.$props.data.text,
            commentID: this.$props.data.commentID,
            timestamp: this.$props.data.timestamp,

            isOwner: false,
        }
    },
    methods: {
        since(timestamp) {
            return utils.since(timestamp);
        },
        goToProfile() {
            this.$router.push(`/profiles/${this.userID}`);
            this.$emit('exit-list-from-entry')
        },
        async deleteComment() {
            try {
                let _ = await this.$axios.delete(`profiles/${this.ownerID}/posts/${this.postID}/comments/${this.commentID}`,
                    { headers: { 'Authorization': `${localStorage.token}` } });
                this.$emit('data-update', {'value': this.commentID, 'opType:': 'delete'});
            } catch (e) {
                this.$emit('error-occurred', e.response.data.message);
            }
        }
    },
    mounted() {
        if (localStorage.userID == this.userID) this.isOwner = true;
    }
}
</script>


<template>
    <div class="comment-entry">
        <div class="comment-entry-info">
            <img class="propic-image" @click="goToProfile" :src="`data:image/jpg;base64,${propic64}`" loading="lazy">
            <span class="profile-entry-username" @click="goToProfile">{{ username }}</span>
            <span class="comment-entry-timestamp">{{ since(timestamp) }}</span>
            <button v-if="!isOwner"  class="comment-entry-button-menu" @click="deleteComment">
                <font-awesome-icon class="comments-icon" icon="fa-regular fa-trash-can" />
            </button>
        </div>
        <span class="comment-entry-text"> {{ text }} </span>
    </div>
</template>


<style>
.comment-entry {
    width: 100%;
    height: auto;
    

    display: flex;
    flex-direction: column;

    border-top: 0.01em solid rgb(0, 0, 0, 0.3);

    padding: 0.5em;
}

.comments-icon {
    font-size: 1em;
    color: rgb(0, 0, 0, 0.6);
}

.comment-entry-info {
    width: 100%;
    height: auto;

    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: left;
}

.propic-image {
    width: 2em;
    height: 2em;
    border-radius: 100%;
    margin-right: 0.5em;

    cursor: pointer;

    border: 0.1em solid rgb(0, 0, 0, 0.1);
}

.profile-entry-username {
    font-size: 1em;
    font-weight: 600;

    cursor: pointer;
}

.comment-entry-button-menu {
    width: 2em;
    height: 2em;
    border: none;
    background-color: white;

    display: flex;
    align-items: center;
    justify-content: center;
}

.comment-entry-text {
    width: 85%;
    height: auto;

    word-wrap: break-word;
    display: inline-block;

    margin-left: 2.8em;

    font-size: 0.8em;
    font-weight: 400;


    color: rgb(0, 0, 0, 0.6);

    position: relative;
}

.comment-entry-timestamp {
    font-size: 0.7em;
    font-weight: 400;
    color: rgb(0, 0, 0, 0.6);

    margin-left: 1em;
}
</style>