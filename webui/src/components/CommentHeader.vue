<script>
export default{
    props: {
        likeCount: { type: Number, required: true },
    },
    data(){
        return{
            isHover: false,
        }
    },
}
</script>


<template>
    <div class="like-header-container">
        <font-awesome-icon icon="fa-regular fa-heart like-icon"/>


    </div>

</template>


<style> 
.like-header-container{
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    width: 100%;
    height: 100%;
    padding: 0 10px;
    box-sizing: border-box;
    cursor: pointer;
}

</style>