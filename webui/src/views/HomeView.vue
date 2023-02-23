<script>
export default {
	data: function () {
		return {
			errormsg: null,
			loading: false,
			posts: [],
		}
	},
	methods: {
		async getMyStream() {
			try {
				let response = await this.$axios.get(`profiles/${localStorage.userID}/feed`, { headers: { 'Authorization': `${localStorage.token}` } });
				this.posts = response.data;
				console.log(this.posts);
			} catch (e) {
				this.errormsg = e.response.data();
			}
		},

		async getMyProfile() {
			this.$router.push(`/profiles/${localStorage.userID}`);
		},
	},

	beforeMount() {
		if (!localStorage.token)
			this.$router.push('/login');
		this.getMyStream();
	},

	mounted() {
	}
}
</script>

<template>
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

	<div class="home-background-centered">
		<div class="floatting-navbar">
			<span @click="getMyProfile" class="left-profile-navbar">
			</span>
			<span class="right-search-navbar">
				<input placeholder="Search..." class="right-searchbar-navbar">
			</span>
		</div>
		<div class="post">
			<Post v-for="post in posts" :postID="post.postID" :userID="post.user.userID" :username="post.user.username"
				:image="post.Image" :caption="post.caption" :timestamp="post.timestamp" :liked="post.liked">
			</Post>
		</div>
	</div>
</template>


<style></style>
