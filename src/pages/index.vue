<template>
    <el-row>
      <el-col :span="8" :gutter="4"></el-col>
      <el-col :span="8" class="mt-4">
        <div class="flex-col justify-center items-center" style="width: 100%;">
          <div @click="handleClick1">AAA</div>
          <div @click="handleClick2">BBB</div>
        </div>
      </el-col>
    </el-row>
  </template>
  
  <script setup>
  import { onMounted, ref } from "vue";
  
  const something = ref(null);
  
  const handleClick1 = () => {
    something.value.onclick1();
  };
  
  const handleClick2 = () => {
    something.value.onclick2();
  };
  
  onMounted(() => {
    something.value = new Something(document.body);
  });
  
  const Something = function (element) {
    // |this| is a newly created object
    this.name = "Something Good";
    this.onclick1 = function (event) {
      console.log(this.name); // undefined, as |this| is the element
    };
  
    this.onclick2 = function (event) {
      console.log(this.name); // 'Something Good', as |this| is bound to newly created object
    };
  
    // bind causes a fixed `this` context to be assigned to onclick2
    this.onclick2 = this.onclick2.bind(this);
  
    element.addEventListener("click", this.onclick1, false);
    element.addEventListener("click", this.onclick2, false); // Trick
  };
  </script>