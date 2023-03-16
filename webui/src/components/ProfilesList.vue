<script>
import CommentFooter from '@/components/CommentFooter.vue';
import CommentEntry from '@/components/CommentEntry.vue';
import SimpleProfileEntry from '@/components/SimpleProfileEntry.vue';
import BanEntry from '@/components/BanProfileEntry.vue';
export default {
    emits: ["exit-list"],
    components: {
        CommentFooter,
        CommentEntry,
        SimpleProfileEntry,
        BanEntry,
    },
    props: {
        dataGetter: { type: Function, required: true },
        dataUpdater: { type: Function, required: false },

        typeList: { type: String, required: true, validator: (value) => ['comment', 'simple', 'ban'].includes(value) },

        textHeader: { type: String, required: true },
        argv: { type: Object, required: false },
    },
    data() {
        return {
            entries: [],
            customEntries: '',

            // Load More
            limit: 10,
            offset: 0,
            busy: false,
            dataAvaible: true,

            additionalData: this.$props.argv,

            errorMsg: '',
            isLoading: false,

        }
    },
    methods: {
        getID(entry) {
            if (this.typeList == 'comment') {
                return entry.commentID;
            } else if (['comment', 'simple', 'ban'].includes(this.typeList)){
                return entry.userID;
            }
        },
        loadMoreContents() {
            if (this.busy || !this.dataAvaible) return;
            this.busy = true;
            this.isLoading = true;
            this.offset += this.limit;
            this.dataGetter(this.entries, this.limit, this.offset, this.dataAvaible);
            this.isLoading = false;
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
    mounted() {
        this.isLoading = true;
        this.dataGetter(this.entries, this.limit, this.offset, this.dataAvaible);
        this.isLoading = false;

        const el = document.getElementsByClassName("list-entries")[0];
        el.addEventListener('scroll', e => {
            if (el.scrollTop + el.clientHeight >= el.scrollHeight) {
                this.loadMoreContents();
            }
        });

        if (this.typeList == 'comment') {
            this.customEntries = CommentEntry;
        } else if (this.typeList == 'simple') {
            this.customEntries = SimpleProfileEntry;
        }
    }
}
</script>


<template>
    <ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"></ErrorMsg>
	<LoadingSpinner :loading=isLoading />


    <div class="list-container-background" @click.self="$emit('exit-list')">
        <div class="list-container">
            <span class="list-header-text">{{ textHeader }}</span>
            <div v-if="typeList == 'simple'" class="list-entries">
                <SimpleProfileEntry class="list-entry" v-for="entry in entries" :key="getID(entry)" :data="entry"
                @exit-list-from-entry="$emit('exit-list')" @error-occured="handleError" />
            </div>
            <div v-else-if="typeList == 'comment'" class="list-entries">
                <CommentEntry class="list-entry" v-for="entry in entries" :key="getID(entry)" :data="entry"
                    @exit-list-from-entry="$emit('exit-list')" @data-update="dataUpdaterEvent"
                    @error-occured="handleError" />
            </div>
            <div v-else class="list-entries">
                <BanEntry class="list-entry" v-for="entry in entries" :key="getID(entry)" :data="entry"
                @exit-list-from-entry="$emit('exit-list')" @error-occured="handleError" />
            </div>

            <span v-if="entries.length == 0" class="empty-list-msg">Nothing to see</span>
            <CommentFooter v-if="typeList == 'comment'" :data="additionalData" @data-update="dataUpdaterEvent"
                @error-occured="handleError" />
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

    display: flex;
    flex-direction: column;
    align-items: center;

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