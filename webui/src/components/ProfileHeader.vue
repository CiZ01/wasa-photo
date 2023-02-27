<script>
import { toRaw } from 'vue';

export default {
    props: {
        userID: { type: Number, required: true },
        username: { type: String, required: true },
        followersCount: { type: Number, required: true },
        followingsCount: { type: Number, required: true },
        bio: { type: String, required: true },
        proPic64: { type: String, required: true },
        isFollowed: { type: Boolean, required: true },
    },
    data() {
        return {
            errorMsg: "",
            isOwner: false,
            textCounter: 0,
            profilesArray: [],
            textHeader: "",

            // Buttons Text
            followText: this.isFollowed ? "Follow" : "Unfollow",
        }
    },
    methods: {
        editingBio() {
            if (this.isOwner) {
                document.getElementsByClassName("top-body-profile-bio-text-counter")[0].style.color = "rgb(0,0,0,0.5)";
                document.getElementsByClassName("top-body-profile-bio-text")[0].style.outline = "auto";
            }
        },
        async saveChangeBio() {
            if (this.isOwner) {
                if (this.bio == "") {
                    this.bio = "This user have notighing to say";
                }
                document.getElementsByClassName("top-body-profile-bio-text")[0].style.outline = "none";
                try {
                    let _ = await this.$axios.put(`/profiles/${this.userID}/bio`, { bio: this.bio }, { headers: { 'Authorization': `${localStorage.token}` } });
                } catch (e) {
                    this.errorMsg = e.toString();
                }
                document.getElementsByClassName("top-body-profile-bio-text-counter")[0].style.color = "#fff";
            }
        },
        countChar() {
            this.textCounter = this.bio.length;
            if (this.bio.includes("\n")) {
                this.bio = this.bio.replace("\n", "");
            }
        },
        editingUsername() {
            if (this.isOwner) {
                document.getElementsByClassName("top-body-profile-username")[0].style.outline = "auto";
            }
        },
        async saveChangeUsername() {
            if (this.isOwner) {
                document.getElementsByClassName("top-body-profile-username")[0].style.outline = "none";
                if (this.username == "" | this.username.length < 3) {
                    this.username = localStorage.username;
                    return
                }
                try {
                    let _ = await this.$axios.put(`/profiles/${this.userID}/username`, { username: this.username }, { headers: { 'Authorization': `${localStorage.token}` } });
                    localStorage.username = this.username;
                } catch (e) {
                    this.errorMsg = e.response.data;
                    this.username = localStorage.username;
                }
            }
        },
        checkUsername() {
            if (this.username.includes("\n")) {
                this.username = this.username.replace("\n", "");
            }
            if (this.username.includes(" ")) {
                this.username = this.username.replace(" ", "");
            }
            if (this.username.includes("\t")) {
                this.username = this.username.replace("\t", "");
            }
            if (this.username.length > 13) {
                this.username = this.username.slice(0, 13);
            }
        },
        async getFollowers() {
            try {
                let response = await this.$axios.get(`/profiles/${this.userID}/followers`, { headers: { 'Authorization': `${localStorage.token}` } });
                this.profilesArray = response.data;
                this.profilesArray = toRaw(this.profilesArray);
                this.textHeader = "Followers";
            } catch (e) {
                this.errorMsg = e.response.data;
            }
        },
        async getFollowings() {
            try {
                let response = await this.$axios.get(`/profiles/${this.userID}/followings`, { headers: { 'Authorization': `${localStorage.token}` } });
                this.profilesArray = response.data;
                this.profilesArray = toRaw(this.profilesArray);
                this.textHeader = "Followings";
            } catch (e) {
                this.errorMsg = e.response.data;
            }
        },
        freeLists() {
            this.profilesArray = [];
        },
        async follow(){
            if (this.isFollowed) {
                try{
                    let _ = await this.$axios.put(`/profiles/${localStorage.userID}/followings/${this.userID}`, { headers: { 'Authorization': `${localStorage.token}` } });
                    this.isFollowed = false;
                    this.followText = "Follow";
                    this.followersCount--;
                } catch (e) {
                    this.errorMsg = e.response.data;
                }
            } else {
                try{
                    let _ = await this.$axios.put(`/profiles/${localStorage.userID}/followings/${this.userID}`, { headers: { 'Authorization': `${localStorage.token}` } });
                    this.isFollowed = true;
                    this.followText = "Unfollow";
                    this.followersCount++;
                } catch (e) {
                    this.errorMsg = e.response.data;
                }
            }
        }
    },
    beforeMount() {
        if (localStorage.userID == this.userID) {
            this.isOwner = true;
        }
    },
    mounted() {
        if (this.isOwner) {
            document.getElementsByClassName("top-body-profile-bio-text")[0].style.cursor = "text";
        }
    }
}
</script>


<template>
    <ErrorMsg v-if="errorMsg" :msg="errorMsg"></ErrorMsg>

    <div class="top-profile-container">
        <div class="top-profile-picture">
            <img :src="`data:image/jpg;base64,${proPic64}`">
        </div>
        <div class="top-body-profile-container">
            <input :readonly="!isOwner" v-model="username" class="top-body-profile-username" @focusin="editingUsername"
                @focusout="saveChangeUsername" @input="checkUsername">
            <div class="top-body-profile-bio-container">
                <span class="top-body-profile-bio-text-counter">{{ textCounter }}/100</span>
                <span class="top-body-profile-bio-label">Bio</span>
                <textarea :readonly="!isOwner" @focusin="editingBio" @focusout="saveChangeBio" @input="countChar"
                    v-model="bio" class="top-body-profile-bio-text" spellcheck="false" maxlength="100" rows="2"></textarea>
            </div>
            <div class="top-body-profile-stats-container">
                <div class="followers-stats" @click="getFollowers">
                    <span class="followers-stats-text">followers</span>
                    <span class="followers-stats-number">{{ followersCount }}</span>
                </div>
                <div class="followers-stats" @click="getFollowings">
                    <span class="followers-stats-text">followings</span>
                    <span class="followers-stats-number">{{ followingsCount }}</span>
                </div>
            </div>

            <div class="top-body-profile-actions">
                <button class="profile-actions-button follow-button" @click="follow()"> {{ followText }} </button>
            </div>
        </div>
    </div>

    <ProfilesList v-if="(profilesArray.length > 0)" :profiles="profilesArray" :textHeader="textHeader"
        :componentEntries="'SimpleProfileEntry'" class="follow-list-view" @exitList="freeLists"> </ProfilesList>
</template>
