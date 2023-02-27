<script>

export default {

	data() {
		return {
			errorMsg: null,
			loading: false,
			posts: [],
			feedLimit: 10,
			feedOffeset: 0,
			busy: false,
		}
	},
	methods: {
		async getMyStream() {
			try {
				const url = `profiles/${localStorage.userID}/feed?limit=${this.feedLimit}&offset=${this.feedOffeset}`;
				let response = await this.$axios.get(url, { headers: { 'Authorization': `${localStorage.token}` } });
				this.posts = response.data;
				console.log(this.posts);
			} catch (e) {
				this.errorMsg = e.response.data;
			}
		},

		getMyProfile() {
			this.$router.push(`/profiles/${localStorage.userID}`);
		},
		updateLike(data){
			this.posts.forEach(post => {
				if (post.postID == data.postID) {
					post.liked = data.liked;
					post.likesCount++;
				}
			});
		}
	},

	beforeMount() {
		if (!localStorage.token) {
			this.$router.push('/login');
			return
		}
		this.getMyStream();
	},

	mounted() {
	}
}
</script>

<template>
	<ErrorMsg v-if="errorMsg" :msg="errorMsg"></ErrorMsg>

	<div class="home-background-centered">
		<div class="floatting-navbar">
			<span @click="getMyProfile" class="left-profile-navbar">
			</span>
			<span class="right-search-navbar">
				<input placeholder="Search..." class="right-searchbar-navbar">
			</span>
		</div>
		<Post v-for="post in posts" :postID="post.postID" :userID="post.user.userID" :username="post.user.username"
			:image="post.image" :caption="post.caption" :timestamp="post.timestamp" :liked="post.liked">
		</Post>
	</div>
</template>