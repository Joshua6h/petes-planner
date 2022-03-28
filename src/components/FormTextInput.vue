<template>
    <v-text-field :label="label" :color="color" :background-color="backgroundColor" outlined :type="isPassword ? 'password' : ''" v-model="contents" @change="updateContents"></v-text-field>
</template>

<script>
export default {
    name: 'FormTextInput',
    data() {
        return {
            contents: ''
        }
    },
    props: {
        label: String,
        isPassword: {
            default: false,
            type: Boolean
        },
        color: {
            type: String,
            default: '#EE4E34'
        },
        backgroundColor: {
            type: String,
            default: 'white'
        }
    },
    methods: {
        updateContents() {
            this.$emit('contentUpdate', this.contents);
        }
    }
}
</script>