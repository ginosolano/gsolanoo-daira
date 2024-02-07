<script>
  import svelteLogo from './assets/svelte.svg'
  import viteLogo from '/vite.svg'
  import Counter from './lib/Counter.svelte'

  
  let inputValue = '';
  let result = '';
  const numbers1 = [7, 8, 9];
  const numbers2 = [4, 5, 6];
  const numbers3 = [3, 2, 1];


  const handleButtonClick = (value) => {
    inputValue += value;
  };

  const handleCalculate = () => {
    try {
      result = eval(inputValue);
    } catch (error) {
      result = 'Error';
    }
  };

  const handleToggleSign = () => {
    if (inputValue && inputValue !== '0') {
      inputValue = inputValue.startsWith('-') ? inputValue.slice(1) : `-${inputValue}`;
    }
  };

  const handleClear = () => {
    inputValue = '';
    result = '';
  };

  const handleClear2 = () => {
    inputValue = inputValue.slice(0, -1);
    result = '';
  };

  const handlePercentage = () => {
    if (inputValue) {
      inputValue = `${parseFloat(inputValue) / 100}`;
    }
  };

</script>



<main>
  <div class="grid grid-cols-2">
    <div class="col-span-1 max-w-md mx-auto my-8 p-6 bg-gray-200 rounded-lg shadow-md">
      <h3 class="mb-4">Calculadora Básica</h3>
      <input
    
      class="bg-black text-white"
      bind:value={inputValue}
      placeholder="0"
      />
       <div class="bg-black text-white mb-4">
        <p class="">{result}</p>
      </div>
   
    
    <div class="grid grid-cols-4 gap-2 mb-4 ">
      <button
        class="col-span-1 bg-orange-500 text-white p-2 rounded"
        on:click={handleClear}
      >
        CE
      </button>
      <button
        class="col-span-1 bg-orange-500 text-white p-2 rounded"
        on:click={handleClear2}
      >
        C
      </button>
      <button
        class="col-span-1 bg-red-500 text-white p-2 rounded"
        on:click={() => handleToggleSign()}
      >
        +/-
      </button>

      <button
        class="col-span-1 bg-red-500 text-white p-2 rounded"
        on:click={() => handlePercentage()}
      >
        %
      </button>
      {#each numbers1 as num (num)}
      <button
        class="col-span-1 bg-blue-500 text-white p-2 rounded"
        on:click={() => handleButtonClick(num)}
      >
      {num}
        </button>
      {/each}
  
      <button
        class="col-span-1 bg-red-500 text-white p-2 rounded"
        on:click={() => handleButtonClick('/')}
      >
      ÷
      </button>
  
      {#each numbers2 as num (num)}
      <button
        class="col-span-1 bg-blue-500 text-white p-2 rounded"
        on:click={() => handleButtonClick(num)}
      >
      {num}
        </button>
      {/each}

      <button
        class="col-span-1 bg-red-500 text-white p-2 rounded"
        on:click={() => handleButtonClick('*')}
      >
        x
      </button>
      
      {#each numbers3 as num (num)}
      <button
        class="col-span-1 bg-blue-500 text-white p-2 rounded"
        on:click={() => handleButtonClick(num)}
      >
      {num}
        </button>
      {/each}

      <button
        class="col-span-1 bg-red-500 text-white p-2 rounded"
        on:click={() => handleButtonClick('-')}
      >
        -
      </button>

      <button
        class="col-span-1 bg-blue-500 text-white p-2 rounded"
        on:click={() => handleButtonClick('0')}
      >
        0
      </button>

      <button
        class="col-span-1 bg-blue-500 text-white p-2 rounded"
        on:click={() => handleButtonClick('.')}
      >
        .
      </button>

      <button
        class="col-span-1 bg-green-500 text-white p-2 rounded"
        on:click={handleCalculate}
      >
        =
      </button>

      <button
      class="col-span-1 bg-red-500 text-white p-2 rounded"
      on:click={() => handleButtonClick('+')}
      >
      +
      </button>
      </div>
      </div>
      <div class="col-span-1 max-w-md mx-auto my-8 p-6 bg-gray-200 rounded-lg shadow-md">
      </div>  
  </div>
  
</main>

<style>
 
</style>
