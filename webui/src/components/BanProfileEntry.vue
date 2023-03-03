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
                    this.$emit('error-occurred',e.toString());
                }
            } else {
                try{
                    let _ = await this.$axios.put(`/profiles/${localStorage.userID}/bans/${this.userID}`,{}, { headers: { "Authorization": `${localStorage.token}` } });
                    this.isBanned = true;
                } catch (e) {
                    this.$emit('error-occurred',e.toString());
                }

            }
        }
    },
}

</script>


<template>
    <div class="profile-entries">
        <img class="propic-image" :src="`data:image/jpg;base64,${propic64}`" loading="lazy">
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
.profile-entries {
    display: flex;
    align-items: center;
    margin-bottom: 0.5em;

    border-top: 0.01em solid rgb(0, 0, 0, 0.3);

    padding: 0.5em 0 0 0.5em;
}

.propic-image{
    width: 2em;
    height: 2em;
    border-radius: 100%;
    margin-right: 0.5em;

    border: 0.1em solid rgb(0, 0, 0, 0.1);
}

.profile-entries-username {
    font-size: 1em;
    font-weight: 500;
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