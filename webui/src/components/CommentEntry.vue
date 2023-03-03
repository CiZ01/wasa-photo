<script>
export default {
    props: {
        comment: {
            type: Object,
            required: true,
        },
    },
    data() {
        return {
            text: this.comment.text,
            username: this.comment.User.username,
            propic64: this.comment.User.userPropicc642
        }
    }
}
</script>

<template>
    <div class="comment-container">
        <img src="">
        <span class="comment-username"></span>
        <span class="comment-text"></span>
    </div>
</template>


<style></style>
