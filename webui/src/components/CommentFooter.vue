<script>
export default {
    emits: ['data-update', 'error-occured'],
    props: {
        data: { type: Object, required: true },
    },
    data() {
        return {
            userID: localStorage.userID,
            username: localStorage.username,
            propic64: localStorage.propic64,

            ownerID: this.data.ownerID,
            postID: this.data.postID,

            commentText: '',
        }
    },
    methods: {
        writingComment() {
            document.getElementsByClassName("comment-input")[0].style.outline = "auto";
            document.getElementsByClassName("comment-input")[0].style.outlineColor = "#03C988";
        },
        noWritingComment() {
            document.getElementsByClassName("comment-input")[0].style.outline = "none";
        },
        async sendComment() {
            if (this.commentText.length == 0) {
                document.getElementsByClassName("comment-input")[0].style.outline = "auto";
                document.getElementsByClassName("comment-input")[0].style.outlineColor = "red";
                return;
            }
            try {
                const body = {
                    text: this.commentText,
                };
                let response = await this.$axios.post(`profiles/${this.ownerID}/posts/${this.postID}/comments`,
                    body,
                    { headers: { 'Authorization': `${localStorage.token}` } });
                this.$emit('data-update', {'value': response.data, 'opType': 'insert'});
                this.commentText = '';
            } catch (e) {
                this.$emit('error-occured', e.toString());
            }
        }
    },
    mounted(){
        console.log(this.data);
    }
}
</script>


<template>
    <div class="comment-form-container">
        <img :src="`data:image/jpg;base64,${propic64}`" class="propic-image">
        <input class="comment-input" type="text" v-model="commentText" placeholder="Write a comment..." spellcheck="false"
            @focus="writingComment" @focusout="noWritingComment" maxlength="64">
        <button class="comment-button-send" type="submit" @click="sendComment">
            <font-awesome-icon class="send-icon" icon="fa-solid fa-paper-plane" />
        </button>
    </div>
</template>


<style>
.comment-form-container {
    width: 100%;
    height: auto;

    position: relative;

    margin: 1em 0 0 0;

    border-top: 1px solid rgb(0, 0, 0, 0.1);

    padding: 1em 0 0 0.5em;

    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: row;
}

.comment-input {
    width: 100%;
    height: 2em;

    border-radius: 0.5em;

    border: none;
    outline: none;

    padding: 0 0.5em 0 0.5em;

    font-size: 1em;
    font-weight: 500;
    color: rgb(0, 0, 0, 0.6);

    margin: 0 0.5em 0 0.5em;
}

.comment-button-send {
    width: 4em;
    height: 2em;

    border-radius: 0.5em;

    border: none;
    outline: none;

    background-color: #03C988;

    font-size: 1em;
    font-weight: 600;
    color: rgb(255, 255, 255, 1);

    margin: 0 0.5em 0 0.5em;
}

.comment-button-send:hover {
    background-color: #03C988;
    opacity: 0.8;
}
</style>