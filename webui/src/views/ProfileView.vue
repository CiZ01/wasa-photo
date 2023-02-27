<script>
import { nextTick } from 'vue';
import PostsGrid from '../components/PostsGrid.vue';
import ProfileHeader from '../components/ProfileHeader.vue'

export default {
    components: {
        PostsGrid,
        ProfileHeader,
    },
    data() {
        return {
            isOwner: false,
            errorMsg: "",

            // Profile data
            userID: parseInt(this.$route.params.userID),
            username: "",
            followersCount: 0,
            followingsCount: 0,
            bio: "This user have notighing to say",
            proPic64: "",
            isFollowed: false,

            // Posts data
            posts:  [],
            offset: 0,
        }
    },
    methods: {
        getProfile() {
            this.$axios.get(`profiles/${this.userID}`, { headers: { 'Authorization': `${localStorage.token}` } }).then(response => {
                this.userID = response.data.user.userID;
                this.username = response.data.user.username;
                this.followersCount = response.data.followersCount;
                this.followingsCount = response.data.followingsCount;
                if (response.data.bio != "")
                    this.bio = response.data.bio;
                this.proPic64 = response.data.user.userPropic64;
                this.isFollowed = response.data.isFollowed;
                console.log(response.data)
            })
                .catch(e => {
                    this.errorMsg = e.toString();
                });
        },
        getPosts() {
            this.posts = [];
            this.$axios.get(`/profiles/${this.userID}/posts?limit=10&offset=${this.offset}`, { headers: { 'Authorization': `${localStorage.token}` } }).then(response => {
                this.posts = response.data;
            }).catch(e => {
                this.errorMsg = e.response.data;
            });
        },
    },
    beforeMount() {
        if (!localStorage.token) {
            this.$router.push('/login');
        }
    },

    mounted() {
        this.getProfile();
        this.getPosts();
        if (this.$route.params.userID == localStorage.userID) {
            this.isOwner = true;
        }
    },

    beforeRouteUpdate(to, from) {
        this.userID = parseInt(to.params.userID);
        this.getProfile();
        this.getPosts();
    }
}

</script>


<template>
    <ProfileHeader :userID="userID" :username="username" :followersCount="followersCount" :followingsCount="followingsCount"
        :bio="bio" :proPic64="proPic64" :isFollowed="isFollowed"> </ProfileHeader>

    <PostsGrid :posts="posts" :profileUserID="parseInt(userID)"> </PostsGrid>

    <div v-if="isOwner" class="delete-account-button-container">
        <button class="delete-account-button"> Delete Account </button>
    </div>
</template>

