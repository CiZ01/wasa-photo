<script>
import SearchBar from "./SearchBar.vue";
export default {
    components:{
        SearchBar,
    },
    props: {

    },
    data() {
        return{

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
    <div class="floatting-navbar">
        <button @click="getMyProfile" class="left-profile-navbar">
        <!-- <img :src="`image/jpg;base64,${localStorage.userID}`"> -->
        </button>
        <div class="upload-post-button"></div>
        <SearchBar> </SearchBar>
    </div>
</template>


<style>
.floatting-navbar {
    background-color: #deefe1;
    height: 5em;
    width: 24em;
    padding: 0;
    border-radius: 10em;

    box-shadow: 0 0.2em 0 0 rgba(176, 179, 184, 0.50);

    position: fixed;
    top: 1em;

    padding: 1.1em;
    margin: 3em;
    
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    z-index: 3;
}

.left-profile-navbar {
    background-color: #ffffff;
    height: 3em;
    width: 3em;
    border-radius: 10em;

    float: left;
    outline: none;
    border: none;

}

.left-profile-navbar:hover {
    cursor: pointer;
}
</style>