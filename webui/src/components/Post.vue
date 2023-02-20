<script>
export default {
    props: ['id', 'username', 'caption', 'image', 'timestamp', 'liked'],
    data(){
        return {
            isHoverComment : false,
            isLiked : this.liked,
        }
    },
    methods:{
        like(){
            this.isLiked = !this.isLiked;
        },
        toggleComment(){
            this.isHoverComment = !this.isHoverComment
        }
    }
}
</script>

<template>
    <div class="post-containter">
        <div class="post-header">
            <img class="post-header-propic-img" src="https://i.imgur.com/2oUpQtj.jpg" alt="" loading="lazy">
            <span class="post-header-username">{{ username }}</span>
        </div>
        <img class="post-img" :src=image @dblclick="like" loading="lazy">
        <div class="post-tail">
            <div class="post-tail-like-comment-options">
                <button class="post-tail-like-button" name="like">
                    <font-awesome-icon v-if=!isLiked icon="fa-regular fa-heart" @click="like"/>
                    <font-awesome-icon v-else icon="fa-solid fa-heart" @click="like" style="color:red"/> 
                </button>
                <button class="post-tail-comment-button" name="comment">
                    <font-awesome-icon v-if="!isHoverComment" icon="fa-regular fa-comment" v-on:mouseover="toggleComment" />
                    <font-awesome-icon v-else icon="fa-solid fa-comment"  v-on:mouseout="toggleComment" style="opacity:0.9"/> 
                </button>
                <button class="post-tail-options-button" name="options">
                    <font-awesome-icon icon="fa-solid fa-ellipsis" />
                </button>
            </div>
            <span class="post-tail-caption" v-if=caption>{{ caption }}</span>
            <span class="post-tail-caption" v-else>Aggiungi una descrizione...</span>
        </div>
    </div>
</template>

<style>
.post-containter {
    width: 30em;
    height: 40em;
    background-color: black;
    border-radius: 1em;
    box-shadow: 0 4px 5px -2px rgba(0, 0, 0, 0.4);
    margin-bottom: 6em;

    display: flex;
    flex-direction: column;
    justify-content: space-between;
}

.post-header {
    width: 100%;
    height: 4em;
    background-color: #fff;
    border-radius: 1em 1em 0 0;
}

.post-header-username {
    font-size: 1.4em;
    font-weight: 600;


    margin-left: 0.5em;

    position: relative;
    top: 0.3em;
}

.post-tail {
    width: 100%;
    height: 8em;
    background-color: #fff;
    border-radius: 0 0 1em 1em;
}

.post-header-propic-img {
    width: 2.5em;
    height: 2.5em;
    border-radius: 10em;

    margin: 0.5em 0em 0em 1em;
}

.post-img {
    width: 100%;
    image-rendering: optimizeSpeed;
}

.post-tail-like-comment-options {
    width: 100%;
    height: 2em;
    padding: 0 0.1em 0 0.1em;
    font-size: 1.5em;
}

.post-tail-like-comment-options >button {
    border: none;
    background-color: transparent;
    outline: none;
    margin: 0.1em 0.3em 0.1em 0.1em;

}

.post-tail-options-button {
    float: right;



}
</style>
