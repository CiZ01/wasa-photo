<script>
export default {
	props: ["loading"]
}
</script>

<template>
	<div v-if="loading" class="loader-background">
		<div style="text-align: center">
			<div class="spinner-border" role="status">
				<span class="visually-hidden">Loading...</span>
			</div>
		</div>
	</div>
	<div v-if="!loading"><slot /></div>
</template>

<style>
.loader-background{
	position: fixed;
	top: 0;
	left: 0;
	width: 100%;
	height: 100%;
	background-color: rgba(255,255,255,0.2);
	z-index: 3;

	display: flex;
	justify-content: center;
	align-items: center;
}
</style>
