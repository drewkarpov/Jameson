<template>
    <span>
        {{ value }} {{ postfix }}
    </span>
</template>


<script>
export default {
  name: 'Description',
  props: ['value'],
  computed: {
    postfix: function () {
      if (this.value === '???') {
        return '';
      }
    }
  },
}

</script>
