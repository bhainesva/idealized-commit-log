import React, { useEffect, useState } from "react";
import * as R from 'ramda';
import {useCombobox} from 'downshift'
import './scss/PackagePicker.scss';
import classNames from "classnames";

interface Props {
  packages: string[],
  onSubmit: (pkg: string) => void
  simple?: boolean,
}

export default function PackagePicker(props: Props) {
  const [filteredPackages, setFilteredPackages] = useState([]);
  const { packages, simple, onSubmit } = props;

  const {
    isOpen,
    inputValue,
    getMenuProps,
    getInputProps,
    getComboboxProps,
    highlightedIndex,
    getItemProps,
  } = useCombobox({
    items: filteredPackages,
    onSelectedItemChange: ({selectedItem}) => {
      onSubmit(selectedItem);
    },
    onInputValueChange: ({inputValue}) => {
      setFilteredPackages(inputValue.length < 2 ? [] : 
        R.take(7, packages.filter(item => item.toLowerCase().includes(inputValue.toLowerCase()))))
    },
  })

  const className = simple ? `PackagePicker PackagePicker--simple` : "PackagePicker";

  const autocompleteOptions = filteredPackages.map((item, index) => (
    <li
      className="Autocomplete-option"
      style={
        highlightedIndex === index ? {backgroundColor: '#bde4ff'} : {}
      }
      key={`${item}${index}`}
      {...getItemProps({item, index})}
    >
      {item}
    </li>
  ))

  return (
    <div className={className} {...getComboboxProps()}>
      {!simple && <h1 className="PackagePicker-label">Choose a package</h1>}
      <form className="PackagePicker-form" onSubmit={(e) => {
          e.preventDefault();
          props.onSubmit(inputValue)}
        }>
          <input type="text" {...getInputProps()} className="PackagePicker-input" />
            <ul className={classNames({'Autocomplete': true, 'u-hidden': !isOpen})} {...getMenuProps()}>
              {isOpen && autocompleteOptions}
            </ul>
        <button className="PackagePicker-submit">Go!</button>
      </form>
    </div>
  )
}