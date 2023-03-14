<script>
import { RectangleStencil, CircleStencil, Cropper } from "vue-advanced-cropper";
import { markRaw } from "vue";
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
        editorType: {
            type: String,
            required: true,
            validator(value) {
                return ['post', 'proPic'].includes(value);
            },
        },
    },
    data() {
        return {
            textCounter: 0,
            textCaption: "",
            imageData: {
                imageFile: null,
                caption: "",
            },

            cropperProps: {
                stencilSize: {
                    width: 0,
                    height: 0,
                },
                stencilComponent: '',
            },
        };
    },
    methods: {
        saveEdit() {

            if (this.editorType === 'post') {
                this.imageData['caption'] = this.textCaption;
            }
            const { canvas, } = this.$refs.cropper.getResult();
            const image64 = canvas.toDataURL("image/jpeg");

            // Converti l'immagine in un blob
            const byteString = atob(image64.split(',')[1]);
            const ab = new ArrayBuffer(byteString.length);
            const ia = new Uint8Array(ab);
            for (let i = 0; i < byteString.length; i++) {
                ia[i] = byteString.charCodeAt(i);
            }
            const blob = new Blob([ab], { type: 'image/jpeg' });

            // Crea l'oggetto file
            const file = new File([blob], 'image/jpeg', { type: 'image/jpeg' });
            this.imageData['imageFile'] = file;
            this.$emit('save-upload-form', this.imageData)
        },
    },
    created() {

        // teoricamente da rimuovere ma voglio sapere se l'errore sussiste
        if (this.editorType === 'post') {
            this.cropperProps.stencilSize = {
                width: 720,
                height: 720,
            };
            this.cropperProps.stencilComponent = markRaw(RectangleStencil);
        } else if (this.editorType === 'proPic') {
            this.cropperProps.stencilSize = {
                width: 250,
                height: 250,
            };
            this.cropperProps.stencilComponent = markRaw(CircleStencil);
        }
    },
    mounted() {
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
        <cropper v-if="editorType == 'post'" class="cropper" :src="image64" :auto-zoom="true" ref="cropper" :stencil-component="RectangleStencil"
            :stencil-size="cropperProps.stencilSize" image-restriction="stencil" :stencil-props="{ aspectRatio: 1 / 1 }" />
        <cropper v-else class="cropper" :src="image64" :auto-zoom="true" ref="cropper" :stencil-component="cropperProps.stencilComponent"
            :stencil-size="cropperProps.stencilSize" image-restriction="stencil" :stencil-props="{ aspectRatio: 1 / 1 }" />
        <div class="caption-form-container" v-if="editorType === 'post'">
            <div class="label-container">
                <span class="caption-title">Caption</span>
                <span class="caption-text-counter">{{ textCounter }}/64</span>
            </div>
            <input v-model="textCaption" class="caption-form" placeholder="There is no caption for this post"
                maxlength="64" />
        </div>
        <div class="button-container">
            <button class="cancel-button" @click="$emit('exit-upload-form')">Cancel</button>
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
  