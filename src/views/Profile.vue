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
                    <h2>Username</h2>
                    <h2>First Last</h2>
                    <h2 class="mb-5">Example.email@go.here</h2>
                    <h3><v-btn icon><v-icon>mdi-border-color</v-icon></v-btn> Edit Information</h3>
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
    }
}
</script>

<style scoped>
.left-side {
    display: flex;
    align-content: flex-start;
}
</style>