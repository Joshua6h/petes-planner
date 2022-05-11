<template>
    <div style="width: 100%"><!--Put sidebar in its own container?-->
        <v-row fluid justify-md="start">
            <v-col md="2" class="pa-10">
                <profile-menu @changePage="setPage($event)"></profile-menu>
            </v-col>
            <v-col v-if="page === 'PROFILE'" md="3"></v-col>
            <v-col v-if="page === 'PROFILE'" md="2"><profile-image></profile-image></v-col>
            <v-col></v-col>
        </v-row>
        <v-row v-if="page === 'PROFILE'">
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
    name: 'ProfilePage',
    data() {
        return {
            page: 'PROFILE'
        }
    },
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
    },
    methods: {
        setPage(page) {
            this.page = page;
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