<template>
    <div>
        <v-row fluid>
            <v-col md="5" class="left-side">
                <v-btn>Friends</v-btn>
                <v-btn>Events</v-btn>
            </v-col>
            <v-col md="2"><profile-image></profile-image></v-col>
            <v-col></v-col>
        </v-row>
        <v-row>
            <v-col md="12">
                <v-container>
                    <h2>Username <v-btn icon><v-icon>mdi-border-color</v-icon></v-btn></h2>
                    <h2>Example.email@go.here</h2>
                </v-container>
                <v-btn depressed class="mt-10">Logout</v-btn>
            </v-col>
        </v-row>
    </div>
</template>

<script>
import ProfileImage from '@/components/ProfileImage.vue';

export default {
    name: 'Profile',
    components: {
        ProfileImage
    }
}
</script>

<style scoped>
.left-side {
    display: flex;
    align-content: flex-start;
}
</style>