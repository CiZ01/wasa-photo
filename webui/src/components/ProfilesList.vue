<script>

export default {
    props: {
        profiles: { type: Object, required: true },
        textHeader: { type: String, required: true },
        componentHeader: { type: Object, required: false },
        level: { type: Number, required: false },
    },
    data() {
        return {
            customLevel: 0,
        }
    },
    methods: {
        exitList() {
            this.$emit("exitList");
        },
        goToProfile(userID) {
            this.$router.push(`/profile/${userID}`);
        },
    },
    beforeMount() {

    },
    mounted() {
    },
}
</script>


<template>
    <div class="list-container-background" @click.self="exitList">
        <div class="list-container">
            <div class="list-container-header">
                <span class="list-header-text">{{ textHeader }}</span>
            </div>
            <div class="list-entries" v-for="profile in profiles">
                <div class="list-entries-img">
                    <img :src="profile.userPropicURL" alt="" loading="lazy">
                </div>
                <span @click="goToProfile(profile.userID)" class="list-entries-username">{{ profile.username }}</span>
            </div>
        </div>
    </div>
</template>

<style>
.list-container-background {
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.2);

    position: fixed;
    top: 0;
    left: 0;
    z-index: 1;

    display: flex;
    justify-content: center;
    align-items: center;
}

.list-container-header{
    width: 100%;
    height: 2em;
    margin-bottom: 1em;

    display: flex;
    align-items: center;
    justify-content: center;
}
.list-header-text {
    font-size: 1.8em;
    font-weight: 600;
    position: relative;


}

.list-container {
    width: 25em;
    height: 30em;

    border-radius: 0.5em;
    background-color: #fff;

    padding: 2em;

    position: relative;
    z-index: 3;

    overflow: scroll;
}

.list-entries {
    width: 100%;
    height: 3em;
    margin-bottom: 0.005em;

    border-top: 0.01em solid rgb(0, 0, 0, 0.3);

    display: flex;
    align-items: center;
    justify-content: left;

    cursor: pointer;
}

.list-entries-img {
    width: 2em;
    height: 2em;
    border-radius: 50%;
    margin-right: 0.5em;

    border: 1px solid black;
}

.list-entries-username {
    font-size: 1em;
    font-weight: 500;
}

.list-container::-webkit-scrollbar {
    display: none;
}

.list-container {
    -ms-overflow-style: none;
    /* IE and Edge */
    scrollbar-width: none;
    /* Firefox */
}
</style>