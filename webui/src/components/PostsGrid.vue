<script>
export default {
    props: {
        posts: { type: Array, required: true },
        profileUserID: { type: Number, required: true },
    },
    data() {
        return {
            errorMsg: "",
            showPost: false,
            postViewData: {},
        }
    },
    methods: {
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
                    post.likesCount++;
                }
            });
        },
        beforeMount() {
        },
        mounted() {
            console.log(this.posts);

        },
    },
}
</script>


<template>
    <div class="posts-grid-container">
        <div v-for="post in posts" :key="post.postID" class="posts-grid-post">
            <img @click="openPost(post)" :src="`data:image/jpg;base64,${post.image}`" loading="lazy"
                class="posts-grid-post-image" :id="post.postID">
        </div>
        <span v-if="(posts.length === 0)" class="posts-grid-nopost-text"> There are no posts yet </span>
    </div>

    <div v-if="showPost" class="post-view" @click.self="exitPost">
        <Post :postID="postViewData.postID" :owner="postViewData.user" :caption="postViewData.caption" :liked="postViewData.liked"
            :timestamp="postViewData.timestamp" :image="postViewData.image" @update-like="updateLike"> </Post>
    </div>
</template>