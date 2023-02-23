<script>
import PostsGrid from '../components/PostsGrid.vue';
import ProfileHeader from '../components/ProfileHeader.vue'

export default {
    data() {
        return {
            isOwner: false,
            errorMsg: "",

            userID: 0,
            username : "",
            followersCount: 0,
            followingsCount: 0,
            bio: "This user have notighing to say",
            proPic: "",
        }
    },
    methods: {
        async getProfile() {
            try {
                let response = await this.$axios.get(`profiles/${this.$route.params.userID}`, { headers: { 'Authorization': `${localStorage.token}` } });
                this.userID = response.data.user.userID;
                this.username = response.data.user.username;
                this.followersCount = response.data.followersCount;
                this.followingsCount = response.data.followingsCount;
                if (response.data.bio != "")
                    this.bio = response.data.bio;
                this.proPic = response.data.user.userPropicURL;
            } catch (e) {
                this.errorMsg = e.toString();
            }
        },
    },
    beforeMount() {
        this.getProfile();
    },

    mounted() {
        if (this.$route.params.userID == localStorage.userID) {
            this.isOwner = true;
        }
    },
}

</script>


<template>
    <ProfileHeader :userID="userID" :username="username" :followersCount="followersCount" :followingsCount="followingsCount" :bio="bio" :proPic="proPic"> </ProfileHeader>

    <PostsGrid> </PostsGrid>

    <div v-if="isOwner" class="delete-account-button-container">
        <button class="delete-account-button"> Delete Account </button>
    </div>
</template>

