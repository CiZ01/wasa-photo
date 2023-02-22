<script>
export default {
	props: ['msg']
}
</script>

<template>
	<div class="alert alert-danger override-style" role="alert">
		{{ msg }}
	</div>
</template>

<style>
.override-style{
	position: absolute;
	top: 2em;
	width: 40em;

	align-items: center;
	}
</style>
