<script>
import SimpleProfileEntry from './SimpleProfileEntry.vue';
export default {
    emits: ['error-occurred'],
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
            busy: false,
            dataAvaible: true,
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
                if (response.data == null) {
                    this.dataAvaible = false;
                    return;
                }
                this.usernameList = [];
                this.usernameList.push(...response.data);
            } catch (e) {
                this.$emit('error-occurred', e.toString());
            }
        },
        exitList() {
            setTimeout(() => {
                this.usernameList = [];
                this.search = "";
                this.dataAvaible = true;
            }, 500);

        },
        loadMoreContents() {
            if (this.busy || !this.dataAvaible) return;
            this.busy = true;
            this.searchOffset += this.searchLimit;
            this.updateSearch();
            this.busy = false;
        },
    },
    mounted() {
        const el = document.querySelectorAll(".users-list")[0];
        el.addEventListener('scroll', e => {
            if (el.scrollTop + el.clientHeight >= el.scrollHeight) {
                this.loadMoreContents();
            }
        });
    },
    watch: {
        search() {
            this.searchOffset = 0;
            this.updateSearch();
        }
    }
}
</script>


<template>
    <div class="search-navbar-container">
        <input autoComplete='none' placeholder="Search..." class="search-bar" v-model="search" @focusout="exitList"
            maxlength="13" spellcheck="false" />
        <div class="users-list" v-show="usernameList.length">
            <SimpleProfileEntry v-for="user in usernameList" :key="user.userID" :data="user"
                @exit-list-from-entry="exitList" />
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

.users-list {
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

    -ms-overflow-style: none;
    /* IE*/
    scrollbar-width: none;
    /* Firefox */
}

.users-list::-webkit-scrollbar {
    display: none;
}
</style>