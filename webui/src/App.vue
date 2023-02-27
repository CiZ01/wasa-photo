<script setup>
import FloatingNavbar from './components/FloattingNavbar.vue'
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	components: {
		FloatingNavbar,
		RouterLink,
		RouterView,
	},
	data() {
		return {
			errorMsg: localStorage.errorMessage,
		}
	},
	methods: {
		goHome() {
			if (this.$route.path == '/home') {
				this.$router.go(0);
				return;
			}
			this.$router.push('/home');
		},
	},
	beforeMount() {
		localStorage.setItem('errorMessage', '');
	},
	mounted() {

	},
}
</script>

<template>
	<main>
		<ErrorMsg v-if="errorMsg" :msg="errorMsg"></ErrorMsg>

		<div class="navbar-header">
			<span @click="goHome">Wasa Photo</span>
		</div>
		<FloatingNavbar v-if="this.$route.path != '/login'" > </FloatingNavbar>

		<RouterView />
	</main>
</template>

<style></style>
