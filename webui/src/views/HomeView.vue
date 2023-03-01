<script>
import FloatingNavbar from '../components/FloattingNavbar.vue'
import UploadPhoto from '../components/UploadPhoto.vue';

export default {
	components: {
		FloatingNavbar,
		UploadPhoto,
	},
	data() {
		return {
			errorMsg: null,
			loading: false,
			posts: [],
			feedLimit: 10,
			feedOffeset: 0,
			busy: false,

			// Search bar
			search: "",
			usernameList: [],
			searchLimit: 10,
			searchOffset: 0,

			// Upload photo
			showUploadPhoto: true,
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
				localStorage.errorMessage =e.response.data;
			}
		},
		updateLike(data){
			this.posts.forEach(post => {
				if (post.postID == data.postID) {
					post.liked = data.liked;
					post.likesCount++;
				}
			});
		},
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

	<UploadPhoto v-if="showUploadPhoto" :photoType="'post'" @exit-upload-form="showUploadPhoto=false"> </UploadPhoto>
	<div class="home-background-centered">
		<FloatingNavbar @show-upload-form="showUploadPhoto=true"> </FloatingNavbar>

		<Post v-for="post in posts" :postID="post.postID" :owner="post.user"
			:image="post.image" :caption="post.caption" :timestamp="post.timestamp" :liked="post.liked">
		</Post>
	</div>
</template>