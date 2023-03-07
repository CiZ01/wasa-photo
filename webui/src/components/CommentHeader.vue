<script>
export default{
    props: {
        data: { type: Object, required: true },
    },
    data(){
        return{
            commentsCount: this.$props.data["commentsCount"],
        }
    },
}
</script>


<template>
    <div class="comment-header-container">
        <font-awesome-icon class="baloon-icon" icon="fa-regular fa-comment"/>
        <span class="comments-count">Comments: {{ commentsCount }}</span>
    </div>

</template>


<style> 
.comment-header-container{
    width: 100%;
    height: auto;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
}

.comments-count{
    font-size: 1.2em;
    font-weight: 600;
    color: rgb(0,0,0,0.6);
}

.baloon-icon{
    font-size: 7em;
    color: rgb(0,0,0,0.3);
}

</style>