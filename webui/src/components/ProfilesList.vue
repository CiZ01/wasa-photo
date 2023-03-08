<script>
import { defineAsyncComponent, shallowRef } from 'vue';


export default {
    emits: ["exit-list"],
    props: {
        dataGetter: { type: Function, required: true },
        dataUpdater: { type: Function, required: false },

        textHeader: { type: String, required: true },
        componentHeader: { type: Object, required: false },
        componentEntries: { type: String, required: true },

        componentFooter: { type: Object, required: false },
        argv: { type: Object, required: false },
    },
    data() {
        return {
            customLevel: 0,
            customEntries: "",
            entries: [],


            // Load More
            limit: 10,
            offset: 0,
            busy: false,
            dataAvaible: true,

            additionalData: this.$props.argv,

            errorMsg: '',

        }
    },
    methods: {
        getID(entry) {
            if (this.componentEntries == 'CommentEntry'){
                return entry.commentID;
            } else if (this.componentEntries == 'SimpleProfileEntry')
                return entry.userID;
        },
        loadMoreContents() {
            if (this.busy || !this.dataAvaible) return;
            console.log("Loading more contents");
            this.busy = true;
            this.offset += this.limit;
            this.dataGetter(this.entries, this.limit, this.offset, this.dataAvaible);
            this.busy = false;

        },
        handleError(msg) {
            this.errorMsg = msg;
        },
        dataUpdaterEvent(values) {
            if (this.dataUpdater) {
                this.dataUpdater(this.entries, values, this.additionalData);
            } else {
                this.errorMsg = "Data Updater not defined";
            }
        },
    },
    beforeMount() {
        this.customEntries = shallowRef(defineAsyncComponent(() =>
            import(`../components/${this.componentEntries}.vue`))
        )
    },
    mounted() {
        this.dataGetter(this.entries, this.limit, this.offset, this.dataAvaible);

        const el = document.getElementsByClassName("list-entries")[0];
        el.addEventListener('scroll', e => {
            if (el.scrollTop + el.clientHeight >= el.scrollHeight) {
                this.loadMoreContents();
            }
        });
    }
}
</script>


<template>
    <ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"></ErrorMsg>


    <div class="list-container-background" @click.self="$emit('exit-list')">
        <div class="list-container">

            <div class="list-entries">
                <div class="list-container-header">
                    <component :is="componentHeader" :data="additionalData" v-if="componentHeader"></component>
                    <span class="list-header-text">{{ textHeader }}</span>
                </div>
                <component class="list-entry" v-for="entry in entries" :key="getID(entry)" :is="customEntries" :data="entry"
                    @exit-list-from-entry="$emit('exit-list')" @data-update="dataUpdaterEvent" @error-occured="handleError">
                </component>

                <span v-if="entries.length == 0" class="empty-list-msg">Nothing to see</span>
            </div>

            <component class="component-footer" :is="componentFooter" v-if="componentFooter" :data="additionalData"
                @data-update="dataUpdaterEvent" @error-occured="handleError"></component>
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
    z-index: 3;

    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}

.list-container-header {
    width: 100%;
    height: auto;

    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
}

.list-header-text {
    font-size: 1.8em;
    font-weight: 600;


}

.list-container {
    width: 25em;
    height: auto;
    min-height: 20em;
    max-height: 40em;

    border-radius: 0.5em;
    background-color: #fff;

    padding: 2em;

    z-index: 2;

    overflow: none;
}

.list-entries {
    width: 100%;
    height: 30em;

    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;

    overflow: scroll;

    -ms-overflow-style: none;
    /* IE and Edge */
    scrollbar-width: none;
    /* Firefox */
}

.list-entry {
    width: 100%;
    height: auto;

    display: flex;
    align-items: center;
    justify-content: left;
}

.list-entries::-webkit-scrollbar {
    display: none;
}

.empty-list-msg {
    font-size: 1.2em;
    font-weight: 600;
    color: #aaa;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
}

.component-footer {
    position: fixed;
}
</style>