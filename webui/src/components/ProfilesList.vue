<script>
import { defineAsyncComponent, shallowRef } from 'vue';


export default {
    emits: ["exit-list"],

    props: {
        profiles: { type: Object, required: true },
        textHeader: { type: String, required: true },
        componentHeader: { type: Object, required: false },
        componentEntries: { type: String, required: true, default: "ProfileEntry" },
        componentFooter: { type: Object, required: false },
        level: { type: Number, required: false },
    },
    data() {
        return {
            customLevel: 0,
            customEntries: "",
        }
    },
    methods: {
    },
    beforeMount() {
        this.customEntries = shallowRef(defineAsyncComponent(() =>
            import(`./${this.componentEntries}.vue`))
        )
    },
}
</script>


<template>
    <div class="list-container-background" @click.self="this.$emit('exit-list')">
        <div class="list-container">
            <div class="list-container-header">
                <span class="list-header-text">{{ textHeader }}</span>
            </div>
            <component class="list-entries" v-for="profile in profiles" :is="customEntries" :data="profile" @exit-list-from-entry="this.$emit('exit-list')">
            </component>
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
    margin-bottom: 0.005em;


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
</style>