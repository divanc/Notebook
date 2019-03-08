import React, { Component } from 'react'
import { connect } from 'react-redux';
import { PropTypes } from 'prop-types';
import {createPost } from '../actions/postActions';

class Postform extends Component {
  constructor(props) {
    super(props);
    this.state = {
      title: '',
      body: ''
    };
  }

  onChange = event => this.setState({[event.target.name]: event.target.value});
 
  onSubmit = event => {
    event.preventDefault();

    const post = {
      title: this.state.title,
      body: this.state.body
    }
    this.props.createPost(post);

  }

  render() {
    return (
      <div>
        <form onSubmit={this.onSubmit}>
        <div>
          <label>Title: </label><br />
          <input type="text" name="title" value={this.state.title} onChange={this.onChange}/>
        </div>
        <div>
          <label> Body: </label><br />
          <input type="text" name="body" value={this.state.body} onChange={this.onChange}/>
        </div>
        <button type='submit'>Submit</button>
        </form>
      </div>
    );
  }
}
Postform.propTypes = {
  createPost: PropTypes.func.isRequired
}

export default connect(null,  { createPost })(Postform);