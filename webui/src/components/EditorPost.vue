<script>
import { Cropper } from "vue-advanced-cropper";
import "vue-advanced-cropper/dist/style.css";

export default {
    emits: ['exit-upload-form', 'save-upload-form'],
    components: {
        Cropper,
    },
    props: {
        image64: {
            type: String,
            required: true,
        },
    },
    data() {
        return {
            textCounter: 0,
            textCaption: "",
            postData: {
                imageBlob: null,
                caption: "",
            },
        };
    },
    methods: {
        saveEdit() {
            const { canvas, } = this.$refs.cropper.getResult();
            this.postData['imageBlob'] = canvas.toDataURL("image/jpeg");
            this.postData['caption'] = this.textCaption;
            this.$emit('save-upload-form', this.postData)
        },
    },
    watch: {
        textCaption() {
            this.textCounter = this.textCaption.length;
        }
    }
};
</script>

<template>
    <div class="cropper-container">
        <cropper class="cropper" :src="image64" :auto-zoom="true" ref="cropper" :stencil-size="{
            width: 720,
            height: 720
        }" image-restriction="stencil" :stencil-props="{
    aspectRatio: 1 / 1,
}"> </cropper>
        <div class="caption-form-container">
            <div class="label-container">
                <span class="caption-title">Caption</span>
                <span class="caption-text-counter">{{ textCounter }}/64</span>
            </div>
            <input v-model="textCaption" class="caption-form" placeholder="There is no caption for this post"
                maxlength="64" />
        </div>
        <div class="button-container">
            <button class="cancel-button" @click="this.$emit('exit-upload-form')">Cancel</button>
            <button class="save-button" @click="saveEdit">Save and Upload</button>
        </div>
    </div>
</template>
  
<style>
.cropper-container {
    background-color: #fff;

    height: auto;
    min-height: 30em;
    max-height: 60em;

    width: auto;
    min-width: 30em;
    max-width: 60em;

    border-radius: 0.5em;

    box-shadow: 0 0.2em 0 0 rgba(176, 179, 184, 0.50);

    position: fixed;

    padding: 2em;

    display: flex;
    flex-direction: column;
    justify-content: space-between;
}

.caption-form-container {
    height: 5em;
    width: auto;

    margin: 1em 0 0 0;

    display: flex;
    flex-direction: column;

}

.label-container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
}

.caption-title {
    font-size: 1em;
    font-weight: 600;
}

.caption-form {
    width: 100%;

    padding: 0.5em;

    border: 0.1em solid rgb(0, 0, 0, 0.9);

    border-radius: 0.5em;

    resize: none;

    font-size: 1em;
    font-weight: 500;
}

.caption-text-counter {
    color: rgb(0, 0, 0, 0.5);

    font-size: 0.8em;
    font-weight: 400;

    float: right;
}

.button-container {
    height: auto;
    width: auto;

    margin-top: 1em;

    display: flex;
    flex-direction: row;
    justify-content: space-around;
}

.button-container button {

    height: 2.5em;
    width: 9em;

    padding: 0;
    border: 0.1em solid rgb(0, 0, 0, 0.9);

    border-radius: 0.5em;

    font-size: 1em;
    font-weight: 600;
}

.button-container button:hover {
    box-shadow: inset 0 0 0 10em rgb(0, 0, 0, 0.1);
}

.save-button {
    background-color: #03C988;
}

.cancel-button {
    background-color: transparent;
}
</style>
  