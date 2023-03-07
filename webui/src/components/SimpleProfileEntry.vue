<script>
export default {
    emits: ["exit-list-from-entry", 'error-occured'],
    props: {
        data: { type: Object, required: true },
    },
    data() {
        return {
            propic64: this.$props.data["userPropic64"],
            username: this.$props.data["username"],
            userID: this.$props.data["userID"],
        }
    },
    methods: {
        goToProfile() {
            this.$router.push(`/profiles/${this.userID}`);
            this.$emit('exit-list-from-entry')
        },
    },
}
</script>


<template>
    <div class="profile-entries" @click="goToProfile">
        <img class="propic-image" :src="`data:image/jpg;base64,${propic64}`" loading="lazy">
        <span class="profile-entries-username">{{ username }}</span>
    </div>
</template>


<style>
.profile-entries {
    display: flex;
    align-items: center;
    margin-bottom: 0.5em;

    cursor: pointer;

    border-top: 0.01em solid rgb(0, 0, 0, 0.3);

    padding: 0.5em 0 0 0.5em;
}

.propic-image{
    width: 2em;
    height: 2em;
    border-radius: 100%;
    margin-right: 0.5em;

    border: 0.1em solid rgb(0, 0, 0, 0.1);
}

.profile-entries-username {
    font-size: 1em;
    font-weight: 600;
}
</style>