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
			errorMsg: '',
			loading: false,
			posts: [],
			feedLimit: 5,
			feedOffeset: 0,

			// Search bar
			search: "",
			usernameList: [],
			searchLimit: 10,
			searchOffset: 0,

			// Upload photo
			showUploadPhoto: false,

			// Load More 
			busy: false,
			dataAvaible: true,
		}
	},
	methods: {
		async getMyStream() {
			try {
				const url = `profiles/${localStorage.userID}/feed?limit=${this.feedLimit}&offset=${this.feedOffeset}`;
				let response = await this.$axios.get(url, { headers: { 'Authorization': `${localStorage.token}` } });
				if (response.data.length == 0) {
					this.dataAvaible = false;
					return;
				}
				this.posts.push(...response.data);
			} catch (e) {
				this.errorMsg = e.toString();
			}
		},
		updateLike(data) {
			this.posts.forEach(post => {
				if (post.postID == data.postID) {
					post.liked = data.liked;
					post.likesCount++;
				}
			});
		},
		loadMoreContents() {
			if (this.busy || !this.dataAvaible) return;
			this.busy = true;
			this.feedOffeset += this.feedLimit;
			this.getMyStream();
			this.busy = false;
		},
		async deletePost(postID) {
			try{
				let _ = await this.$axios.delete(`profiles/${localStorage.userID}/posts/${postID}`, { headers: { 'Authorization': `${localStorage.token}` } });
				this.posts = this.posts.filter(post => post.postID != postID);
			}catch(e){
				this.errorMsg = e.toString();
			}
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

		document.addEventListener('scroll', e => {
			if (document.documentElement.scrollTop + window.innerHeight >= document.documentElement.scrollHeight * (0.7)) {
				this.loadMoreContents();
			}
		});
	}
}
</script>

<template>
	<ErrorMsg v-if="errorMsg" :msg="errorMsg"></ErrorMsg>

	<UploadPhoto v-if="showUploadPhoto" :photoType="'post'" @exit-upload-form="showUploadPhoto = false"
		@refresh-data="getMyStream()" @error-occured="errorMsg = value"> </UploadPhoto>
	<FloatingNavbar @show-upload-form="showUploadPhoto = true"> </FloatingNavbar>

	<Post v-for="post in posts" :postID="post.postID" :owner="post.user" :image="post.image" :caption="post.caption"
		:timestamp="post.timestamp" :liked="post.liked" @delete-post="deletePost" />
</template>