<script>
export default {
	emits: ['close-error'],
	props: ['msg']
}
</script>

<template>
	<div class="alert alert-danger override-style" role="alert">
		{{ msg }}
		<button @click="$emit('close-error')">
			<font-awesome-icon icon="fa-solid fa-times" />
		</button>
	</div>
</template>

<style>
.override-style {
	position: absolute;
	top: 2em;

	min-width: 40em;
	width: auto;
	max-width: 70em;

	height: auto;

	align-items: center;

	z-index: 5;

	display: flex;
	flex-direction: row;
	justify-content: space-between;
	align-items: center;

}

.override-style button {
	font-size: small;

	margin-left: 1em;
	background-color: transparent;
	border: none;
	outline: none;
}

.override-style button:hover {
	background-color: rgb(0, 0, 0, 0.1);
}
</style>
