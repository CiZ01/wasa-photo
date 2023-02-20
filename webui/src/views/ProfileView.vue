<script>
export default {
    data() {
        return {
            userID: this.$route.params.userID,
            username: "",
            followerCount: 0,
            followingCount: 0,
            errorMsg: "",
        }
    },
    methods: {
        async getProfile() {
            try {
                let response = await this.$axios.get(`profiles/${this.userID}`, { headers: { 'Authorization': `${localStorage.token}` } });
                this.username = response.data.username;
                this.followerCount = response.data.followerCount;
                this.followingCount = response.data.followingCount;
            } catch (e) {
                this.errorMsg = e.toString();
            }
        },
    },
    beforeMount() {
        this.getProfile();
    },
}
</script>


<template>
    <div class="top-profile-container">
        <div class="top-profile-picture">
            <img src="" alt="">
        </div>
        <div class="top-body-profile-container">
            <span class="top-body-profile-username">francesco000000</span>
            <div class="top-body-profile-bio-container">
                <span class="top-body-profile-bio">bio</span>
            </div>
            <div class="top-body-profile-stats-container">
                <div class="followers-stats">
                    <span class="followers-stats-text">followers</span>
                    <span class="followers-stats-number">0</span>
                </div>
                <div class="followers-stats">
                    <span class="followers-stats-text">following</span>
                    <span class="followers-stats-number">0</span>
                </div>
            </div>
        </div>
    </div>
</template>
