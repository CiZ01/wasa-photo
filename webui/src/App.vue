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
		<div class="navbar-header">
			<span @click="goHome">Wasa Photo</span>
		</div>
		<RouterView />
	</main>
</template>

<style></style>
