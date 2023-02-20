<script>
export default {
	data: function () {
		return {
			errormsg: null,
			loading: false,
			posts: [ {postID: '200', imageURL: "http://localhost:3000/wasa_photo/storage/1/posts/6.jpg", liked: false, user: {username: "pippo"}}],
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/");
				this.some_data = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async getMyStream() {
			try {
				let response = await this.$axios.get(`profiles/${localStorage.userID}/feed`, { headers: { 'Authorization': `${localStorage.token}` } });
				this.posts = response.data;
				console.log(this.posts);

			} catch (e) {
				this.errormsg = e.toString();
				console.log(this.errormsg);
			}
		},

		async getMyProfile(){
			this.$router.push(`/profiles/${localStorage.userID}`);
		},

		mounted() {
			this.getMyStream();
			this.refresh();
		}
	}
}
</script>

<!--sfondo grigio da social newtork che conterrà i posts-->
<template>
	<div class="bg_pattern Polka">
		<div class="floatting-navbar">
			<span @click="getMyProfile" class="left-profile-navbar">
			</span>
			<span class="right-search-navbar">
				<input placeholder="Search..." class="right-searchbar-navbar">
			</span>
		</div>
		<div class="home-background-centered">
			<div class="post">
				<Post v-for="post in posts" :key="post.postID" :username="post.user.username" :image="post.imageURL"
					:caption="post.caption" :liked="post.liked"></Post>
			</div>
		</div>
	</div>
<!--ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg--></template>


<style></style>
