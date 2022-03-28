<template>
    <v-img src="@/assets/profilePlaceholder.png" @click="show"></v-img>
</template>

<script>
export default {
    name: 'ProfileImage',
    methods: {
        show() {
            console.log("Image clicked");
        }
    }
}
</script>