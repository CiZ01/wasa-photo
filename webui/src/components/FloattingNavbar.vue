<script>
import SearchBar from "./SearchBar.vue";
export default {
    components: {
        SearchBar,
    },
    emits: ['show-upload-form'],
    props: {

    },
    data() {
        return {
            propic64: localStorage.propic64,
            errorMsg: "",
        }
    },
    methods: {
        getMyProfile() {
            this.$router.push(`/profiles/${localStorage.userID}`);
        },
    },
}
</script>

<template>
    <ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"></ErrorMsg>
    <div class="floatting-navbar">
        <button @click="getMyProfile" class="left-profile-navbar">
            <img :src="`data:image/jpg;base64,${propic64}`">
        </button>
        <button class="home-button">
            <font-awesome-icon icon="fa-sharp fa-regular fa-house-blank" />
        </button>
        <button class="upload-post-button" @click="$emit('show-upload-form')">
            <font-awesome-icon icon="fa-regular fa-square-plus" />
        </button>
        <SearchBar> </SearchBar>
    </div>
</template>


<style>
.floatting-navbar {
    background-color: #deefe1;
    height: auto;
    width: auto;
    padding: 0;
    border-radius: 10em;

    box-shadow: 0 0.2em 0 0 rgba(176, 179, 184, 0.50);

    position: sticky;
    top: 1em;

    padding: 1em;
    margin-bottom: 1em;

    display: flex;
    flex-direction: row;
    justify-content: space-between;
    z-index: 3;
}

.left-profile-navbar img {
    height: 3em;
    width: 3em;
    border-radius: 10em;
    overflow: hidden;

}

.left-profile-navbar {
    background-color: rgb(0, 0, 0, 0.7);
    height: 3em;
    width: 3em;
    border-radius: 10em;

    padding: 0;

    float: left;
    outline: none;
    border: none;

}

.left-profile-navbar:hover {
    cursor: pointer;
}

.upload-post-button {
    background-color: rgb(255, 255, 255, 0.4);
    height: 1.5em;
    width: 1.5em;
    border-radius: 10em;

    padding: 0;

    margin: 0 0.5em 0 0.5em;

    float: left;
    outline: none;
    border: none;

    font-size: 2em;
    line-height: 0;

    color: rgb(0, 0, 0, 0.8)
}

.upload-post-button:hover {
    box-shadow: inset 0 0 0 1em rgb(0, 0, 0, 0.2);
}

.home-button {
    background-color: rgb(255, 255, 255, 0.4);
    height: 1.5em;
    width: 1.5em;
    border-radius: 10em;

    padding: 0;
    margin: 0 0 0 0.5em;

    float: left;
    outline: none;
    border: none;

    font-size: 2em;
    line-height: 0;

    color: rgb(0, 0, 0, 0.8)
}
</style>