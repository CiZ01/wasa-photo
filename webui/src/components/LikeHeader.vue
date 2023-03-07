<script>
export default{
    props: {
        data: { type: Object, required: true },
    },
    data(){
        return{
            isHover: false,
            likesCount: this.$props.data["likesCount"],
        }
    },
}
</script>


<template> 
    <div class="like-header-container">
        <font-awesome-icon class="heart-icon" icon="fa-regular fa-heart"/>
        <span class="likes-count">Likes: {{ likesCount }}</span>
    </div>
</template>


<style> 
.like-header-container{
    width: 100%;
    height: auto;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
}

.likes-count{
    font-size: 1.2em;
    font-weight: 600;
    color: rgb(0,0,0,0.6);
}

.heart-icon{
    font-size: 7em;
    color: rgb(0,0,0,0.3);
}
</style>