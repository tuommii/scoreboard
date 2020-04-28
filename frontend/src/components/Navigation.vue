<template>
  <div class="navbar is-fixed-bottom has-text-centered">

    <div class="columns is-mobile">
      <div class="column left"><a class="button bold" @click="navigate(showPrev)">
        Prev <span class="tag">{{showPrev}}</span></a>
      </div>
      <div class="column middle active is-2"><a class="button">{{active}}</a></div>
      <div class="column right"><a class="button bold" @click="navigate(showNext)">
        Next <span class="tag">{{showNext}}</span></a></div>
    </div>
</div>
</template>

<script>
export default {
  props: ['active', 'basketCount'],
  data() {
    return {
    }
  },
  methods: {
    navigate(num) {
      this.$emit('navigate', num);
    }
  },
  computed: {
    showNext() {
      if (this.active >= this.basketCount)
        return (1);
      return this.active + 1;
    },
    showPrev() {
      if (this.active <= 1)
        return (this.basketCount);
      return this.active - 1;
    }
  }
};
</script>

<style scoped>
a, a:visited {
  color: #fff !important;
}
.button {
  border: 0;
  background: inherit;
}

.bold {
  font-weight: 700;
}

.column {
  border: 1px solid #24444e;
  background: #1f2743;
}
.left {
  border-right: 1px solid;
}
.right {
  border-left: 1px solid;
}
.tag {
  color: #ccc;
  background: #000;
  margin-left: 5px;
}
</style>
