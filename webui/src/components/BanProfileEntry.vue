<script>
export default {
    emits: ['exit-list-from-entry', 'error-occurred'],
    props: {
        data: { type: Object, required: true },
    },
    data() {
        return {
            propic64: this.$props.data["userPropic64"],
            username: this.$props.data["username"],
            userID: this.$props.data["userID"],
            isBanned: true,
        }
    },
    methods: {
        goToProfile() {
            this.$router.push(`/profiles/${this.userID}`);
            this.$emit('exit-list-from-entry')
        },
        async ban(){
            if (this.isBanned){
                try{
                    let _ = await this.$axios.delete(`/profiles/${localStorage.userID}/bans/${this.userID}`, { headers: { "Authorization": `${localStorage.token}` } });
                    this.isBanned = false;
                } catch (e) {
                    this.$emit('error-occurred',this.$utils.errorToString(e));
                }
            } else {
                try{
                    let _ = await this.$axios.put(`/profiles/${localStorage.userID}/bans/${this.userID}`,{}, { headers: { "Authorization": `${localStorage.token}` } });
                    this.isBanned = true;
                } catch (e) {
                    this.$emit('error-occurred',this.$utils.errorToString(e));
                }

            }
        }
    },
}

</script>


<template>
    <div class="profile-entries no-click">
        <img class="propic-image no-click" :src="`data:image/jpg;base64,${propic64}`" loading="lazy">
        <span class="profile-entries-username">{{ username }}</span>
        <div class="ban-button-container">
            <button class="ban-button" @click="ban">
                <span v-if="isBanned">Unban</span>
                <span v-else>Ban</span>
            </button>
        </div>
    </div>
</template>


<style>
.no-click{
    cursor: default;
}
.ban-button-container{
    margin-left: auto;
}

.ban-button{
    background-color: rgb(0, 0, 0, 0);
    border: 0.1em solid rgb(0, 0, 0, 0.3);
    border-radius: 0.5em;
    padding: 0.2em 0.5em;
    font-size: 0.8em;
    font-weight: 500;
    color: rgb(0, 0, 0, 0.3);
}

.ban-button:hover{
    background-color: rgb(0, 0, 0, 0.1);
    color: rgb(0, 0, 0, 0.5);

}
</style>