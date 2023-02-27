<script>
export default {
    emits: ["exit-list-from-entry"],
    props: {
        data: { type: Object, required: true },
    },
    data() {
        return {
            propic64: this.$props.data["propic64"],
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
    <div class="profile-entries" @click.native="goToProfile">
        <div class="profile-entries-propic">
            <img :src="propic64" alt="" loading="lazy">
        </div>
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
}

.profile-entries-propic {
    width: 2em;
    height: 2em;
    border-radius: 50%;
    margin-right: 0.5em;

    border: 1px solid black;
}

.profile-entries-username {
    font-size: 1em;
    font-weight: 500;
}
</style>