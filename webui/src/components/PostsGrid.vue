<script>
export default {
    data() {
        return {
            posts: [],
            profileUserID: this.$route.params.userID,
            offset: 0,
            url: "",
            errorMsg: "",
            showPost: false,
            postViewData: {},
        }
    },
    methods: {
        async getPosts() {
            try {
                this.posts = [];
                let response = await this.$axios.get(`/profiles/${this.profileUserID}/posts?limit=10&offset=${this.offset}`, { headers: { 'Authorization': `${localStorage.token}` } });
                this.posts = response.data;
            } catch (e) {
                this.errorMsg = e.response.data;
            }
        },
        imageSrc() {
            return new URL(`../../../../storage/1/posts/6.jpg`, import.meta.url).href;
        },
        openPost(post) {
            this.showPost = true;
            this.postViewData = post;
        },
        exitPost() {
            this.showPost = false;
            this.postViewData = {};
        },
    },
    beforeMount() {
        this.getPosts();
    },
    mounted() {
    },
}
</script>


<template>
    <ErrorMsg v-if="errorMsg" :msg="errorMsg"></ErrorMsg>
    <div class="posts-grid-container">
        <div v-for="post in posts" :key="post.postID" class="posts-grid-post">
            <img @click="openPost(post)" :src="imageSrc" loading="lazy" class="posts-grid-post-image" :id="post.postID">
        </div>
        <span v-if="(posts.length === 0)" class="posts-grid-nopost-text"> There are no posts yet </span>
    </div>

    <div v-if="showPost" class="post-view" @click.self="exitPost"> 
            <Post :postID="postViewData.postID" :userID="postViewData.user.userID" :username="postViewData.user.username" :image="postViewData.imageURL"
            :caption="postViewData.caption" :liked="postViewData.liked"> </Post>
    </div>

</template>