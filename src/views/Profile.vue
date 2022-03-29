<template>
    <div>
        <v-row fluid>
            <v-col md="5"></v-col>
            <v-col md="2"><profile-image></profile-image></v-col>
            <v-col></v-col>
        </v-row>
        <v-row>
            <v-col md="5"></v-col>
            <v-col md="2">
                <h2>Username</h2>
                <v-btn depressed class="mt-10">Logout</v-btn>
            </v-col>
            <v-col></v-col>
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