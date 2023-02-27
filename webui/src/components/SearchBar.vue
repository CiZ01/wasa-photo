<script>
import SimpleProfileEntry from './SimpleProfileEntry.vue';
export default {
    components: {
        SimpleProfileEntry,
    },
    props: {

    },
    data() {
        return {
            search: "",
            usernameList: [],
            searchLimit: 10,
            searchOffset: 0,
        }
    },
    methods: {
        async updateSearch() {
            if (this.search.length == 0) {
                this.usernameList = [];
                return;
            }
            try {
                const url = `profiles?limit=${this.searchLimit}&offset=${this.searchOffset}&search=${this.search}`;
                let response = await this.$axios.get(url, { headers: { 'Authorization': `${localStorage.token}` } });
                this.usernameList = response.data;
            } catch (e) {
                localStorage.errorMessage =
 e.response.data;
            }
        },
        exitList(){
            setTimeout(() => {
                this.usernameList = [];
                this.search = "";
            }, 100);

        }
    },
}
</script>


<template>
    <div class="search-navbar-container">
        <input placeholder="Search..." class="search-bar" @input="updateSearch" v-model="search" @focusout="exitList">
        <div class="users-list" v-if="usernameList.length"> 
            <SimpleProfileEntry v-for="user in usernameList" :key="user.userID" :data="user" @exit-list-from-entry="exitList">
            </SimpleProfileEntry>
        </div>
    </div>
</template>


<style>
.search-navbar-container {
    background-color: #ffffff;
    height: 3em;
    width: 12em;
    border-radius: 10em;

    float: right;
}

.search-bar {
    background-color: #f0f2f5;
    height: 3em;
    width: 12em;
    border-radius: 10em;
    padding: 1em;

    box-shadow: none;
    outline: none;
    border: none;
}

.users-list{
    background-color: #ffffff;
    height: auto;
    max-height: 20em;

    overflow: auto;

    width: 12em;
    border-radius: 0 0 0.5em 0.5em;
    padding: 2em 1em 0.5em 1em;

    position: relative;
    top: -1em;
    z-index: -1;

    box-shadow: 0px 0.5em 0.6em -0.5em #000000;
    outline: none;
    border: none;
}
</style>