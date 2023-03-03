<script>
export default {
    props: {
        msg: { type: String, required: true },
        type: { type: String, required: true, validator(value) {
            return ['dialog-box', 'success'].includes(value)
        }},
        successFunction: { type: Function, required: false },
        cancelFunction: { type: Function, required: false },
    },
    data(){
        return {
            
        }
    }
}
</script>


<template>
<div class="popup-container">
    <div class="msg-container">
        <span class="msg">{{ msg }}</span>
    </div>
    <div class="buttons-container">
        <button v-if="type == 'dialog-box'" class="cancel-button" @click="cancelFunction">Cancel</button>
        <button class="confirm-button" @click="successFunction">Confirm</button>
    </div>
</div>
</template>


<style>
</style>