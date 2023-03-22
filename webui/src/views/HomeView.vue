<script>
import FloatingNavbar from '@/components/FloattingNavbar.vue'
import UploadPhoto from '@/components/UploadPhoto.vue';
export default {
	components: {
		FloatingNavbar,
		UploadPhoto,
	},
	data() {
		return {
			errorMsg: "",
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

			isLoading: false,
		}
	},
	methods: {
		getID(post) {
			return `${post.user.userID}` + `${post.postID}`;
		},
		async getMyStream() {
			this.isLoading = true;
			try {
				const url = `profiles/${localStorage.userID}/feed?limit=${this.feedLimit}&offset=${this.feedOffeset}`;
				let response = await this.$axios.get(url, { headers: { 'Authorization': `${localStorage.token}` } });
				if (response.data == null) {
					this.dataAvaible = false;
					this.isLoading = false;
					return;
				}
				this.posts.push(...response.data);
			} catch (e) {
				this.errorMsg = this.$utils.errorToString(e);;
			}
			this.isLoading = false;
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
			const index = this.posts.findIndex(post => post.postID == postID && post.user.userID == localStorage.userID);
			try {
				await this.$axios.delete(`profiles/${localStorage.userID}/posts/${postID}`, { headers: { 'Authorization': `${localStorage.token}` } });
				this.posts.splice(index, 1);
			} catch (e) {
				this.errorMsg = $utils.errorToString();
			}
		},
	},
	mounted() {
		if (!localStorage.token) {
			this.$router.push('/login');
			return
		}
		this.getMyStream();

		document.addEventListener('scroll', e => {
			if (document.documentElement.scrollTop + window.innerHeight >= document.documentElement.scrollHeight * (0.7)) {
				this.loadMoreContents();
			}
		});
	},
}
</script>

<template>
	<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"></ErrorMsg>

	<LoadingSpinner :loading=isLoading />

	<UploadPhoto v-if="showUploadPhoto" :photoType="'post'" @exit-upload-form="showUploadPhoto = false"
		@refresh-data="$router.go(0)" @error-occurred="(value) => { errorMsg = value }" />
	<FloatingNavbar @show-upload-form="showUploadPhoto = true" @error-occurred="(value) => { errorMsg = value }" />

	<span v-if="posts.length == 0" class="no-posts-text"> There are no posts yet </span>
	<span v-if="posts.length == 0" class="no-posts-text fw-500 fs-6"> Start to follow someone!</span>
	<Post v-for="post in posts" :key="getID(post)" :postData="post" @delete-post="deletePost"
		@error-occurred="(value) => { errorMsg = value }" />
</template>