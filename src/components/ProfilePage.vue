<template>
    <div style="width: 100%">
        <v-row fluid justify-md="start">
            <v-col md="2" class="pa-10">
                <profile-menu @changePage="setPage($event)"></profile-menu>
            </v-col>
            <v-col md="8">
                <profile-card v-if="page === 'PROFILE'"></profile-card>
                <friend-card v-if="page === 'FRIENDS'"></friend-card>
            </v-col>
        </v-row>
    </div>
</template>

<script>
import ProfileMenu from '@/components/ProfileMenu.vue';
import ProfileCard from '@/components/ProfileCard.vue';
import FriendCard from '@/components/FriendCard.vue';

export default {
    name: 'ProfilePage',
    components: {
        ProfileMenu,
        ProfileCard,
        FriendCard
    },
    data() {
        return {
            page: 'PROFILE'
        }
    },
    methods: {
        setPage(page) {
            this.page = page;
        }
    }
}
</script>