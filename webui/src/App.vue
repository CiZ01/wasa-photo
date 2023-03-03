<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	components: {
		RouterLink,
		RouterView,
	},
	data() {
		return {
			errorMsg: '',
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
		<RouterView />
	</main>
</template>

<style></style>
