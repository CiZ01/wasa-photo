<script>
import EditorPost from '../components/EditorPost.vue';
export default {
    components: {
        EditorPost,
    },
    emits: ['exit-upload-form', 'refresh-data', 'error-occurred'],
    props: {
        photoType: {
            type: String,
            required: true,
            validator(value) {
                return ['post', 'proPic'].includes(value)
            },
        },
    },
    data() {
        return {
            file64: "",

            errorMsg: '',
        }
    },
    methods: {
        onChange() {
            const file = this.$refs.file.files[0];

            // Check file type, only jpg and jpeg are allowed
            const fileType = file.type;
            if (fileType !== "image/jpeg") {
                this.errorMsg = "File type not supported, only jpg and jpeg are allowed";
                document.querySelector('.drag-drop-area').style.backgroundColor = "#FF8989";
                return
            }

            // Check file size, max 5MB
            const fileSize = file.size;
            if (fileSize > 5242880) {
                console.log("cioa");
                this.errorMsg = "File size is too big. Max size is 5MB";
                document.querySelector('.drag-drop-area').style.backgroundColor = "#FF8989";
                return
            }

            // Convert file to base64
            this.file64 = URL.createObjectURL(file);
        },

        saveData(data) {
            if (this.photoType == 'post') {
                this.createPost(data);
            } else if (this.photoType == 'proPic') {
                this.changePropic(data);
            }
        },

        async createPost(postData) {
            const formData = new FormData();
            formData.append("image", postData['imageFile']);
            formData.append("caption", postData['caption']);

            try {
                let _ = await this.$axios.post(`profiles/${localStorage.userID}/posts`, formData, {
                    headers: {
                        'Authorization': `${localStorage.token}`,
                        'content-type': 'multipart/form-data'
                    }
                });
                this.$emit('refresh-data')
                setTimeout(this.$emit('exit-upload-form'), 1000);
            } catch (e) {
                this.errorMsg = this.$utils.errorToString(e);;
                this.$emit('error-occurred', this.errorMsg);
            }
        },

        async changePropic(propicData) {
            const formData = new FormData();
            formData.append("image", propicData['imageFile']);
            try {
                let _ = await this.$axios.put(`profiles/${localStorage.userID}/profile-picture`, formData, {
                    headers: {
                        'Authorization': `${localStorage.token}`,
                        'content-type': 'multipart/form-data'
                    }
                });
                this.$emit('refresh-data');
                this.$emit('exit-upload-form');
            } catch (e) {
                this.errorMsg = this.$utils.errorToString(e);;
                this.$emit('error-occurred', this.errorMsg);
            }
        }
    },
    watch: {
        errorMsg(){
            this.$emit('error-occurred', this.errorMsg);
            console.log(this.errorMsg)
        },
    },
}
</script>


<template>
    <div class="upload-form-background" @click.self="this.$emit('exit-upload-form')">
        <div class="upload-form-container" v-if="!file64">
            <div class="drag-drop-area-container">
                <button class="drag-drop-area" @click="this.$refs.file.click()">
                    <input type="file" ref="file" accept=".jpg,.jpeg" @change="onChange" hidden />
                    <span class="drag-drop-area-text">
                        Drop your photo here
                    </span>
                    <span class="drag-drop-area-subtext">
                        max size 5MB, only jpg, jpeg
                    </span>

                </button>
            </div>
            <div class="bottom-area">
                <button @click="this.$refs.file.click()" class="upload-button">Choose File
                    <input type="file" ref="file" accept=".jpg,.jpeg" @change="onChange" hidden />
                </button>

            </div>
        </div>

        <EditorPost :image64="file64" :editorType="this.$props.photoType" v-if="file64"
            @exit-upload-form="this.$emit('exit-upload-form')" @save-upload-form="saveData" />
    </div>
</template>


<style>
.upload-form-background {
    background-color: rgba(0, 0, 0, 0.5);
    height: 100%;
    width: 100%;
    padding: 0;
    z-index: 4;

    position: fixed;
    top: 0;
    left: 0;

    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}

.upload-form-container {
    background-color: #fff;
    height: 40em;
    width: 50em;
    padding: 0;
    border-radius: 0.5em;

    box-shadow: 0 0.2em 0 0 rgba(176, 179, 184, 0.50);

    position: fixed;

    padding: 2em;

    display: flex;
    flex-direction: column;
    justify-content: space-between;

}

.drag-drop-area-container {

    height: 100%;
    width: 100%;

    padding: 0;


    cursor: pointer;
}

.drag-drop-area {
    background-color: #deefe1;

    height: 100%;
    width: 100%;

    border-radius: 0.5em;
    border: 0.3em dashed rgb(0, 0, 0, 0.5);

    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-between;


    position: relative;

    opacity: 1;
}

.drag-drop-area-text {
    font-size: 1.5em;
    font-weight: 500;
    color: rgb(0, 0, 0, 0.5);
    text-align: center;
    margin-top: 9em;
}

.drag-drop-area-subtext {
    font-size: 1em;
    font-weight: 500;
    color: rgb(0, 0, 0, 0.5);
    text-align: center;
    margin-bottom: 15em;
}

.bottom-area {
    height: 5em;
    width: auto;

    padding: 1em;
    display: flex;
    flex-direction: row;
    justify-content: center;

}

.upload-button {
    background-color: #03C988;

    height: 2.5em;
    width: 9em;

    padding: 0;

    border-radius: 10em;
    border: none;

    font-size: 1.2em;
    font-weight: 600;
}
</style>