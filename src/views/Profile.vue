<template>
    <div>
        <v-row fluid>
            <v-col md="5" class="left-side">
                <profile-menu></profile-menu>
            </v-col>
            <v-col md="2"><profile-image></profile-image></v-col>
            <v-col></v-col>
        </v-row>
        <v-row>
            <v-col md="12">
                <v-container>
                    <h2>{{username}}</h2>
                    <h2>{{firstname}} {{lastname}}</h2>
                    <h2 class="mb-5">{{email}}</h2>
                    <!-- <h3><v-btn icon><v-icon>mdi-border-color</v-icon></v-btn> Edit Information</h3> -->
                    <!-- We probably don't want them to be able to change their info at the moment -->
                </v-container>
            </v-col>
        </v-row>
    </div>
</template>

<script>
import ProfileImage from '@/components/ProfileImage.vue';
import ProfileMenu from '@/components/ProfileMenu.vue';

export default {
    name: 'Profile',
    components: {
        ProfileImage,
        ProfileMenu
    },
    computed: {
        username() {
            return this.$store.getters.username;
        },
        firstname() {
            return this.$store.getters.firstname;
        },
        lastname() {
            return this.$store.getters.lastname;
        },
        email() {
            return this.$store.getters.email;
        }
    }
}
</script>

<style scoped>
.left-side {
    display: flex;
    align-content: flex-start;
}
</style>