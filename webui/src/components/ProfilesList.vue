<script>
import { defineAsyncComponent, shallowRef } from 'vue';


export default {
    emits: ["exit-list"],

    props: {
        dataGetter: { type: Function, required: true },
        textHeader: { type: String, required: true },
        componentHeader: { type: Object, required: false },
        componentEntries: { type: String, required: true},
        typeEntry: { type: String, required: false, default: 'followings', validator(value) {
            return ['bans', 'followings'].includes(value)
        } },
        componentFooter: { type: Object, required: false },
        level: { type: Number, required: false },
    },
    data() {
        return {
            customLevel: 0,
            customEntries: "",
            profiles: [],


            // Load More
            limit: 10,
            offset: 0,
            busy: false,
            dataAvaible: true,

            errorMsg: '',

        }
    },
    methods: {
        loadMoreContents(){
            if (this.busy || !this.dataAvaible) return;
            this.busy = true;
            this.offset += this.limit;
            this.dataGetter(this.profiles, this.limit, this.offset, this.dataAvaible);
            this.busy = false;
            
        },
        handleError(msg){
            this.errorMsg = msg;
        }
    },
    beforeMount() {
        this.customEntries = shallowRef(defineAsyncComponent(() =>
            import(`./${this.componentEntries}.vue`))
        )
    },
    mounted(){
        this.dataGetter(this.profiles, this.limit, this.offset, this.dataAvaible);

        
        const el = document.getElementsByClassName("list-container")[0];
        el.addEventListener('scroll', e => {
            if(el.scrollTop + el.clientHeight >= el.scrollHeight) {
                this.loadMoreContents();
            }
        });
    }
}
</script>


<template>

    <ErrorMsg v-if="errorMsg" :msg="errorMsg"></ErrorMsg>


    <div class="list-container-background" @click.self="this.$emit('exit-list')">
        <div class="list-container">
            <div class="list-container-header">
                <span class="list-header-text">{{ textHeader }}</span>
            </div>
            <component class="list-entries" v-for="user in profiles" :key="user.userID" :is="customEntries" :data="user" :type="typeEntry" 
            @exit-list-from-entry="this.$emit('exit-list')" @error-occured="handleError">
            </component>

            <span v-if="profiles.length == 0" class="empty-list-msg">Nothing to see</span>
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

.list-container-header {
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
    z-index: 2;

    overflow: scroll;
}

.list-entries {
    width: 100%;
    height: 3em;

    display: flex;
    align-items: center;
    justify-content: left;
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

.empty-list-msg{
    font-size: 1.2em;
    font-weight: 600;
    color: #aaa;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
}
</style>